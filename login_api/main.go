package main

/*
import (
	//ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	model "login-api/src/model"
)

func init() {}

func main() {
	/*
		msgChan := make(chan *ckafka.Message)
		consumer := kafka_config.NewConsumer(msgChan)
		go consumer.Consumer()

		for msg := range msgChan {
			go kafka_config.SendProduce(msg)
			println(string(msg.Value))
		}

	user := model.User{
		email:     "teste",
		username:  "test",
		password:  "test",
		isAdmin:   true,
		isNewUser: false,
	}
	result := model.NewResult()

	model.User.
		c.ValidateUser(user, result)

}*/

import "fmt"

type Address struct {
	city    string
	country string
}

type User struct {
	name    string
	age     int
	address Address
}

func main() {

	p := User{
		name: "John Doe",
		age:  34,
		address: Address{
			city:    "New York",
			country: "USA",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("Country:", p.address.country)
}
