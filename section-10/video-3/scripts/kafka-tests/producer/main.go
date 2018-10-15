package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

type sagaState uint16

const (
	SAGA_START sagaState = iota
	BOUGHT_VIDEO_START
	BOUGHT_VIDEO_END
	UPDATE_USER_ACCOUNT_START
	UPDATE_USER_ACCOUNT_END
	UPDATE_MANAGER_ACCOUNT_START
	UPATE_MANAGER_ACCOUNT_END
	SAGA_END

	SAGA_ROLLBACK_START
	BOUGHT_VIDEO_ROLLBACK_START
	BOUGHT_VIDEO_ROLLBACK_END
	UPDATE_USER_ACCOUNT_ROLLBACK_START
	UPDATE_USER_ACCOUNT_ROLLBACK_END
	UPDATE_MANAGER_ACCOUNT_ROLLBACK_START
	UPATE_MANAGER_ACCOUNT_ROLLBACK_END
	SAGA_ROLLBACK_END
)

type sagaMessage struct {
	State   sagaState `json:"saga_state"`
	UserID  uint32    `json:"user_id"`
	VideoID uint32    `json:"video_id"`
	Offset  uint64    `json:"offset"`
}

func main() {
	// to produce messages
	topic := "test"

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// brokers := []string{"192.168.59.103:9092"}
	brokers := []string{"localhost:9092"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		// Should not reach here
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			// Should not reach here
			panic(err)
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("Hola!"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
}
