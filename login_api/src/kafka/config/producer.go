package kafka_config

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"login-api/src/controller/criptografia"
	"login-api/src/model"
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

func SendProduce(msg *ckafka.Message) {

	producer := newProducer()
	user := model.NewUser()

	json.Unmarshal(msg.Value, &user)

	// call user validation
	result := model.NewResult()

	result = criptografia.ValidateUser(user, result)

	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	err := enc.Encode(result)
	if err != nil {
		log.Printf("Encode error: %v", err)
	}

	Produce(network.Bytes(), producer)
}

//retorna error sempre que não consegue enviar msg
func Produce(r []byte, producer *ckafka.Producer) error {

	topic := os.Getenv("topics")

	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          r,
	}

	err := producer.Produce(message, nil)
	if err != nil {
		return err
	}

	return nil
}
