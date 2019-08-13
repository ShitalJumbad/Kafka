package main

import (
	"fmt"
	

	"github.com/Shopify/sarama"
)

const topic = "newest"
var brokers = []string{"localhost:9092"}

func newProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

func prepareMessage(topic, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}

	return msg
}

func main() {
	producer, err := newProducer()
	if err != nil {
		fmt.Println("Could not create producer: ", err)
	}

	msg := prepareMessage(topic, "Shital Jumbad")
	partition, offset, err := producer.SendMessage(msg)
	fmt.Println(partition)
	fmt.Println(offset)
	fmt.Println("Done printing")
	
	
}
