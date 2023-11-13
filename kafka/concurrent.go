package main

import (
	"context"
	"encoding/json"
	"hash/fnv"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/qiniu/x/log"
)

type MsgInfo struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	Job    string `json:"job"`
}

type KafkaClientConfig struct {
	Brokers []string `json:"brokers"`
	Version string   `json:"version"`
}

type ConsumerConfig struct {
	KafkaClientConfig
	Topics        []string `json:"topics"`
	GroupId       string   `json:"group_id"`
	ConcurrentNum int      `json:"concurrent_num"` // 每个消费者实例内部的并发数
}

type KafkaConsumer struct {
	consumerConfig     *ConsumerConfig
	consumerGroup      []sarama.ConsumerGroup
	bizConsumerHandler []*Consumer
}

func main() {
	kcc := KafkaClientConfig{
		Brokers: []string{"100.100.142.91:9092"},
	}

	cc := &ConsumerConfig{
		KafkaClientConfig: kcc,
		Topics:            []string{"lzf_1"},
		GroupId:           "lizhenfeng",
		ConcurrentNum:     2,
	}

	waitErr := make(chan error, 1)
	ctx, cancel := context.WithCancel(context.Background())

	consumerInstance, err := NewKafkaConsumer(cc)
	if err != nil {
		log.Error("new kafka client err:", err)
	}

	go func() {
		err := consumerInstance.Start(ctx)
		if err != nil {
			waitErr <- err
			log.Error("consumerInstance error:", err)
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-waitErr:
		log.Error("terminating:", err)
	case <-sigterm:
		log.Info("terminating: via signal")
	}
	cancel()
}

func NewKafkaConsumer(config *ConsumerConfig) (*KafkaConsumer, error) {
	var err error
	cc := sarama.NewConfig()
	cc.Consumer.Return.Errors = true
	cc.Consumer.Offsets.Initial = sarama.OffsetOldest
	cAdmin, err := sarama.NewClusterAdmin(config.Brokers, cc)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// 根据partition并发
	topicMetas, err := cAdmin.DescribeTopics(config.Topics)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	partitionCount := 0
	for _, meta := range topicMetas {
		for _, p := range meta.Partitions {
			if p.Err == sarama.ErrNoError {
				partitionCount += 1
			}
		}
	}
	log.Infof("Got partition count: %d", partitionCount)

	concurrentNum := partitionCount
	consumerGroupList := []sarama.ConsumerGroup{}
	handlerList := []*Consumer{}

	c := &KafkaConsumer{
		consumerConfig: config,
	}

	for i := 0; i < concurrentNum; i++ {
		// 内部并行处理函数
		concurrentExcutor := NewConcurrentExcutor(config.ConcurrentNum, c.externalExcute)
		consumerGroup, err := sarama.NewConsumerGroup(config.Brokers, config.GroupId, cc)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		consumerGroupList = append(consumerGroupList, consumerGroup)
		bizConsumer := &Consumer{number: i, excutor: concurrentExcutor}
		handlerList = append(handlerList, bizConsumer)
	}
	log.Infof("Successfully init %v consumer instances of consumer group %s", concurrentNum, config.GroupId)

	c.consumerGroup = consumerGroupList
	c.bizConsumerHandler = handlerList

	return c, nil
}

func (c *KafkaConsumer) externalExcute(ctx context.Context, msg MsgInfo) error {
	bs, err := json.Marshal(msg)
	if err != nil {
		log.Error(err)
	}
	log.Infof("consume message: %s", string(bs))
	return nil
}

func (c *KafkaConsumer) Start(ctx context.Context) error {
	wg := sync.WaitGroup{}
	for i := 0; i < len(c.consumerGroup); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			e := c.bizConsumerHandler[i].excutor
			e.start()            // 执行消息
			err := c.Run(ctx, i) // 消费消息
			if err != nil {
				log.Error(err)
			}
		}(i)
	}
	wg.Wait()
	log.Info("all the consumer goroutines exited")
	return nil
}

func (c *KafkaConsumer) Run(ctx context.Context, i int) error {
	for {
		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		if err := c.consumerGroup[i].Consume(ctx, c.consumerConfig.Topics, c.bizConsumerHandler[i]); err != nil {
			log.Error("Error from consumer:", err)
			return err
		}
		// check if context was cancelled, signaling that the consumer should stop
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	// ready chan bool
	number  int // 表示第几个consumer实例
	excutor *ConcurrentExcutor
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	// close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	session.Commit()
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/IBM/sarama/blob/main/consumer_group.go#L27-L29
	ctx := context.Background()
	for message := range claim.Messages() {
		log.Infof("got message: timestamp = %v, topic = %s, partition = %d, offset = %d, handler = %d", message.Timestamp, message.Topic, message.Partition, message.Offset, consumer.number)
		if len(message.Value) == 0 {
			continue
		}
		var msgInfo MsgInfo
		err := json.Unmarshal(message.Value, &msgInfo)
		if err != nil {
			log.Error("json unmarshal err:", err)
		}
		consumer.excutor.PushChan(ctx, msgInfo)
		session.MarkMessage(message, "")
	}

	return nil
}

type ConcurrentExcutor struct {
	// 并行数
	Num int `json:"num"`
	// 锁
	Mux sync.RWMutex `json:"mux"`
	// 通道
	Chans map[int]chan interface{}
	// 外部执行逻辑的函数
	ExLogicFunc ExternalLogicFunc
}

func NewConcurrentExcutor(num int, exLogicFunc ExternalLogicFunc) *ConcurrentExcutor {
	if num == 0 {
		num = 10
	}
	if exLogicFunc == nil {
		log.Panic("ConcurrentExcutor ExternalLogicFunc nil")
	}
	chans := make(map[int]chan interface{}, num)
	for i := 0; i < num; i++ {
		chans[i] = make(chan interface{}, 1)
	}
	ce := &ConcurrentExcutor{
		Num:         num,
		Chans:       chans,
		ExLogicFunc: exLogicFunc,
	}
	return ce
}

// 外部执行逻辑的函数
type ExternalLogicFunc func(ctx context.Context, msg MsgInfo) error

// 推送到通道
func (ce *ConcurrentExcutor) PushChan(ctx context.Context, msg MsgInfo) {
	key := msg.Job
	mod := ce.getMod(key)
	ce.Mux.RLock()
	defer ce.Mux.RUnlock()
	ch, ok := ce.Chans[mod]
	if !ok {
		log.Panicf("ConcurrentExcutor PushChan err, key = %s, mod = %d", key, mod)
	}
	ch <- msg
	log.Infof("message sent to channel %d", mod)
}

// 取模算法
func (ce *ConcurrentExcutor) getMod(key string) int {
	return int(hash(key)) % ce.Num
}

// 开始并行计算
func (ce *ConcurrentExcutor) start() error {
	for i := 0; i < ce.Num; i++ {
		go func(i int) {
			ch := ce.Chans[i]
			for data := range ch {
				msg := data.(MsgInfo)
				ce.ExLogicFunc(context.Background(), msg)
			}
		}(i)
	}
	return nil
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
