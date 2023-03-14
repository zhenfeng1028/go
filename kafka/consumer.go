package main

import (
	"context"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/qiniu/x/log"
)

/*
	测试步骤：
	1.进入kafka目录（/usr/local/Cellar/kafka/3.0.0）
	2.在终端开启zookeeper：bin/zookeeper-server-start libexec/config/zookeeper.properties
	3.在新的终端开启kafka：bin/kafka-server-start libexec/config/server.properties
	4.执行consumer.go：go run consumer.go
	5.在新的终端生产消息：bin/kafka-console-producer --topic test-topic --bootstrap-server localhost:9092
	6.在执行该文件的终端可以看到生产的消息被消费了
*/

func main() {
	var ctx = context.Background()

	kcc := KafkaClientConfig{
		Brokers: make([]string, 0),
		Version: "2.3.0",
	}
	kcc.Brokers = append(kcc.Brokers, "localhost:9092")

	cc := ConsumerConfig{
		KafkaClientConfig: kcc,
		Topics:            "test-topic",
		Group:             "test",
	}
	c, err := NewKafkaConsumer(cc)
	if err != nil {
		log.Fatalf("NewKafkaConsumer err: %v", err)
	}

	go func() {
		defer c.Close()
		err := c.Run(ctx)
		if err != nil {
			log.Errorf("consumer stopped : %v", err)
			return
		}
	}()

	for {
		data := <-c.Consume()
		msg, ok := data.([]byte)
		if !ok {
			log.Warnf("data is not of type []byte")

		}
		log.Printf("receive data: %v\n", string(msg))
	}
}

type KafkaClientConfig struct {
	Brokers []string `json:"brokers"`
	Version string   `json:"version"`
}

type ConsumerConfig struct {
	KafkaClientConfig
	Topics string `json:"topics"`
	Group  string `json:"group"`
}

type Data interface{}

// KafkaConsumer represents a Sarama consumer group consumer
type KafkaConsumer struct {
	ch     chan Data
	config *ConsumerConfig
	client sarama.ConsumerGroup
}

type Consumer interface {
	Run(ctx context.Context) error
	Close() error
	Consume() <-chan Data
}

func NewKafkaConsumer(c ConsumerConfig) (Consumer, error) {
	version, err := sarama.ParseKafkaVersion(c.Version)
	if err != nil {
		log.Errorf("Error parsing Kafka version: %s", c.Version)
		return nil, err
	}

	/**
	* Construct a new Sarama configuration.
	* The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	client, err := sarama.NewConsumerGroup(c.Brokers, c.Group, config)
	if err != nil {
		log.Error("Error create consumer group : ", c.Brokers, c.Group, config)
		return nil, err
	}

	csm := &KafkaConsumer{
		config: &c,
		client: client,
		ch:     make(chan Data, 10),
	}

	return csm, nil
}

func (c *KafkaConsumer) Close() error {
	return c.client.Close()
}

func (c *KafkaConsumer) Consume() <-chan Data {
	return c.ch
}

func (c *KafkaConsumer) Run(ctx context.Context) error {
	log.Println("begin to run consumer , topics:", c.config.Topics)
	for {
		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		if err := c.client.Consume(ctx, strings.Split(c.config.Topics, ","), c); err != nil {
			log.Errorf("Error from consumer: %v", err)
			return err
		}
		// check if context was cancelled, signaling that the consumer should stop
		if ctx.Err() != nil {
			log.Error("ctx error:", ctx.Err())
			return ctx.Err()
		}
		// c.ready = make(chan bool)
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *KafkaConsumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	// close(c.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *KafkaConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (c *KafkaConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		c.ch <- message.Value
		session.MarkMessage(message, "")
	}

	return nil
}
