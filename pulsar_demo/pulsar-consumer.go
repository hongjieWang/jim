package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"time"
)

func main() {
	fmt.Println("Pulsar Consumer")
	//实例化Pulsar client
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://110.40.141.168:6650",
	})

	if err != nil {
		log.Fatal(err)
	}

	//使用client对象实例化consumer
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "my-topic",
		SubscriptionName: "sub-demo",
		Type:             pulsar.Shared,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close()

	ctx := context.Background()

	//无限循环监听topic
	for {
		msg, err := consumer.Receive(ctx)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Received message : %v  %s \n", string(msg.Payload()), time.Now().Format("2006-01-02 15:04:05"))
		}

		consumer.Ack(msg)

	}

}
