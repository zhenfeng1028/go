package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

// kafka consumer

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Println("fail to start consumer,", err)
		return
	}
	partitionList, err := consumer.Partitions("quickstart-events") // 根据topic取到所有的分区
	if err != nil {
		fmt.Println("fail to get list of partition,", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("quickstart-events", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d, err: %v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition: %d Offset: %d Key: %v Value: %v\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
	select {}
}
