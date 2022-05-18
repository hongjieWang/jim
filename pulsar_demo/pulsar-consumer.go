package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"time"
)

const token = "eyJrZXlJZCI6InB1bHNhci13NDVqeno1MjN2Ym4iLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJwdWxzYXItdzQ1anp6NTIzdmJuX2ppbSJ9.i0xRiumZs5HN9-xtgp5j4WtHEbxzalqWVD1JmuVYph4"
const topic = "pulsar-w45jzz523vbn/jim/message"
const URL = "http://pulsar-w45jzz523vbn.tdmq.ap-bj.public.tencenttdmq.com:8080"

func main() {
	fmt.Println("Pulsar Consumer")
	//实例化Pulsar client
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:            URL,
		Authentication: pulsar.NewAuthenticationToken(token),
	})

	if err != nil {
		log.Fatal(err)
	}

	//使用client对象实例化consumer
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
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
