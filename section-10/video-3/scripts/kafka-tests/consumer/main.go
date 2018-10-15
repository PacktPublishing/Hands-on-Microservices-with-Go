package main

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
)

func main() {

	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	// Specify brokers address. This is default one
	brokers := []string{"localhost:2181"}

	// Create new consumer
	consumer, err := consumergroup.JoinConsumerGroup("test1", []string{"test"}, brokers, config)
	if err != nil {
		panic(err)
	}

	for msg := range consumer.Messages() {
		log.Println(msg)
		consumer.CommitUpto(msg)
	}

}
