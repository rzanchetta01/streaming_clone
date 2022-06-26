package kafka_config

import (
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func newProducer() *ckafka.Producer {
	config := &ckafka.ConfigMap{
		"bootstrap.server": os.Getenv(""),
	}

	producer, err := ckafka.NewProducer(config)
	if err != nil {
		log.Fatalln("KAFKA PRODUCER CREATIN ERROR : {}", err.Error())
	}

	return producer
}

//retorna error sempre que não consegue enviar msg
func Produce(isReal byte, topic string, producer *ckafka.Producer) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte{isReal},
	}

	err := producer.Produce(message, nil)
	if err != nil {
		return err
	}

	return nil
}
