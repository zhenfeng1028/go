package main

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/qiniu/x/log"
)

// 使用KRaft模式（无需Zookeeper）
/* 创建docker-compose-kraft.yml文件，内容如下：
version: '3'
services:
kafka:
	image: bitnami/kafka:latest
	ports:
	- "9092:9092"
	environment:
	- KAFKA_ENABLE_KRAFT=yes
	- KAFKA_CFG_PROCESS_ROLES=broker,controller
	- KAFKA_CFG_NODE_ID=1
	- KAFKA_CFG_CONTROLLER_QUORUM_VOTERS=1@kafka:9093
	- KAFKA_CFG_LISTENERS=PLAINTEXT://:9092,CONTROLLER://:9093
	- KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092
	- KAFKA_CFG_CONTROLLER_LISTENER_NAMES=CONTROLLER
	volumes:
	- kafka_data:/bitnami
volumes:  # 必须显式声明volumes
kafka_data:
	driver: local
*/

func main() {
	ctx := context.Background()

	kcc := KafkaClientConfig{
		Brokers: make([]string, 0),
		Version: "2.3.0",
	}
	kcc.Brokers = append(kcc.Brokers, "localhost:9092")

	cc := ConsumerConfig{
		KafkaClientConfig: kcc,
		Topics:            []string{"test-topic"},
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
			log.Error("consumer stopped:", err)
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
	Topics []string `json:"topics"`
	Group  string   `json:"group"`
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
		log.Error("Error create consumer group:", c.Brokers, c.Group, config)
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
	log.Info("begin to run consumer, topics:", c.config.Topics)
	for {
		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		if err := c.client.Consume(ctx, c.config.Topics, c); err != nil {
			log.Error("Error from consumer:", err)
			return err
		}
		// check if context was cancelled, signaling that the consumer should stop
		if ctx.Err() != nil {
			log.Error("ctx error:", ctx.Err())
			return ctx.Err()
		}
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *KafkaConsumer) Setup(sarama.ConsumerGroupSession) error {
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
	// https://github.com/IBM/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		log.Printf("Message claimed: timestamp = %v, topic = %s", message.Timestamp, message.Topic)
		c.ch <- message.Value
		session.MarkMessage(message, "")
	}

	return nil
}
