package main

import (
	"log"
	kafka_config "login-api/src/kafka/config"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERROR READING .ENV FILE")
	}
}

func main() {

	msgChan := make(chan *ckafka.Message)
	consumer := kafka_config.NewConsumer(msgChan)
	go consumer.Consumer()

	for msg := range msgChan {
		go kafka_config.SendProduce(msg)
		println(string(msg.Value))
	}

}
