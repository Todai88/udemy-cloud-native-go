package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	fmt.Println("Starting synchronous Kafka subscriber...")
	time.Sleep(5 * time.Second)

	// create the consumer configuration
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.CommitInterval = 5 * time.Second
	config.Consumer.Return.Errors = true

	// Create new consumer for given brokers and configuration
	brokers := []string{brokerAddr()}
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	// decide about the offset here: literal value, sarama.OffsetOlders, sarama.OffsetNewest
	// this is important in case of reconnection
	// create a kafka partition consumer for given topic
	consumer, err := master.ConsumePartition(topic(), 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Count how many messages have been processed
	msgCount := 0

	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			// Consume and process the message(s)
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Println("Recieved messages", string(msg.Key), string(msg.Value))
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
}

func brokerAddr() string {
	brokerAddr := os.Getenv("BROKER_ADDR")
	if len(brokerAddr) == 0 {
		brokerAddr = "localhost:9092"
	}
	return brokerAddr
}

func topic() string {
	topic := os.Getenv("TOPIC")
	if len(topic) == 0 {
		topic = "default-topic"
	}
	return topic
}
