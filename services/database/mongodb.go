package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var client = &mongo.Client{}
var UserAccountCollection = &mongo.Collection{}
var MessageTemplateCollection = &mongo.Collection{}
var BusinessCollection = &mongo.Collection{}

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://julywhj:125846whj@cluster0.r1o1v.mongodb.net/test"))
	// mongodb数据表
	const collection = "message_template"
	const business = "business"
	const userAccount = "user_account"
	UserAccountCollection = client.Database("jim").Collection(userAccount)
	MessageTemplateCollection = client.Database("jim").Collection(collection)
	BusinessCollection = client.Database("jim").Collection(business)
}
