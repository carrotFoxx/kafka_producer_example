package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type host_msg struct {
	Id   int
	Task int
	Arg  string
}

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func main() {
	// get kafka writer using environment variables.
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("topic")
	writer := newKafkaWriter(kafkaURL, topic)
	defer writer.Close()
	fmt.Println("start producing ... !!")
	for i := 0; ; i++ {
		hm := host_msg{
			Id:   2027925481,
			Task: 55,
			Arg:  "is test message",
		}
		bts, _ := json.Marshal(hm)
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("Key-%d", i)),
			Value: bts,
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}
