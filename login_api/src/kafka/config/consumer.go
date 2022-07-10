package kafka_config

import (
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	Msg chan *ckafka.Message
}

func NewConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		Msg: msgChan,
	}
}

func (kafka *KafkaConsumer) Consumer() {
	config := &ckafka.ConfigMap{
		"bootstrap.server": os.Getenv("bootstrap.servers"),
		"group.id":         os.Getenv("group.id"),
	}

	consumer, err := ckafka.NewConsumer(config)
	if err != nil {
		log.Fatalf("ERROR CONSUMING KAFKA MESSAGE")
	}

	topics := []string{os.Getenv("topics")}
	consumer.SubscribeTopics(topics, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			kafka.Msg <- msg
		}
	}
}
