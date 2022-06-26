package main

import (
	"fmt"
	"login-api/src/controller/criptografia"
	kafka_config "login-api/src/kafka/config"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func init() {}

func main() {

	msgChan := make(chan *ckafka.Message)
	consumer := kafka_config.NewConsumer(msgChan)
	go consumer.Consumer()

	/*for msg := range msgChan {
		go kafka_service.Produce(msg)
		println(string(msg.Value))
	}*/

	var c = criptografia.CriptografiaMD5{
		Texto: "Teste",
	}

	fmt.Println(c.Criptografa())

	for {

	}

}
