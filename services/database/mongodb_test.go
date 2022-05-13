package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
)

var client *mongo.Client
var database string
var collection string

func init() {
	client, _ = Init("mongodb+srv://julywhj:XXXX@cluster0.r1o1v.mongodb.net/test")
	database = "jim"
	collection = "jim_notice_config"
}

func Benchmark_ping(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
			panic(err)
		}
		fmt.Println("Successfully connected and pinged.")
	})
}

func Benchmark_install(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(1024)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		c := client.Database(database).Collection(collection)
		doc := bson.D{{"title", "Invisible Cities"}, {"author", "Italo Calvino"}, {"year_published", 1974}}
		result, _ := c.InsertOne(context.TODO(), doc)
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	})
}
