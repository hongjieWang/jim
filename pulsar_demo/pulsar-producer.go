package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"time"
)

func main() {
	fmt.Println("Pulsar Producer")

	ctx := context.Background()

	//实例化Pulsar client
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://110.40.141.168:6650",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})

	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client:%v", err)
	}

	// 创建producer
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic",
	})

	if err != nil {
		log.Fatalf("Could not instantiate Pulsar producer:%v", err)
	}

	defer producer.Close()

	msg := &pulsar.ProducerMessage{
		Payload:      []byte("Hello,This is a message from Pulsar Producer!"),
		DeliverAfter: 10 * time.Second,
	}

	if err, _ := producer.Send(ctx, msg); err != nil {
		log.Fatalf("Producer could not send message:%v", err)
	}

}
