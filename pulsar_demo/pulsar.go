package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"time"
)

func main() {

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://110.40.141.168:6650",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	topicName := "jim"
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic:           topicName,
		DisableBatching: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topicName,
		SubscriptionName: "subName",
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	ID, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload:      []byte(fmt.Sprintf("test")),
		DeliverAfter: 3 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ID)

	ctx, canc := context.WithTimeout(context.Background(), 1*time.Second)
	msg, err := consumer.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg.Payload())
	canc()

	ctx, canc = context.WithTimeout(context.Background(), 5*time.Second)
	msg, err = consumer.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg.Payload())
	canc()
}
