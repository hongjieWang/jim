package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var clientTest *mongo.Client

// Database 测试数据库
const Database = "jim"

// Collection 测试Collection
const Collection = "user"

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientTest, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://julywhj:XXXX@cluster0.r1o1v.mongodb.net/test"))
}

// Ping 验证mongodb链接成功
func Ping() {
	if err := clientTest.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
}

// DemoUser 测试用户struct
type DemoUser struct {
	Name  string `bson:"name" json:"name,omitempty"`
	Phone string `bson:"phone" json:"phone,omitempty"`
	Age   int64  `bson:"age" json:"age,omitempty"`
}

// InstallOne 插入一行数据
func InstallOne(client *mongo.Client) {
	collection := client.Database(Database).Collection(Collection)
	one, err := collection.InsertOne(context.TODO(), DemoUser{
		Name:  "张三",
		Phone: "18346566786",
		Age:   28,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(one)
}

// InstallMany 批量插入
func InstallMany(client *mongo.Client) {
	collection := client.Database(Database).Collection(Collection)
	users := []interface{}{
		DemoUser{Name: "汪鸣晨", Phone: "13659630001", Age: 28},
		DemoUser{Name: "盖流惠", Phone: "13659630004", Age: 28},
		DemoUser{Name: "桓野云", Phone: "13659630004", Age: 28},
		DemoUser{Name: "郝语彤", Phone: "18345654901", Age: 28},
		DemoUser{Name: "沈振文", Phone: "18345654902", Age: 28},
		DemoUser{Name: "古雪翎", Phone: "18345654903", Age: 28},
		DemoUser{Name: "燕丁辰", Phone: "18345654904", Age: 28},
		DemoUser{Name: "乔任真", Phone: "18345654905", Age: 28},
		DemoUser{Name: "暴笑妍", Phone: "18345654906", Age: 28},
		DemoUser{Name: "池雁风", Phone: "18345654907", Age: 28},
		DemoUser{Name: "马玉萍", Phone: "18345654908", Age: 28},
		DemoUser{Name: "崔子舒", Phone: "18345654909", Age: 28},
		DemoUser{Name: "简晗晗", Phone: "18345654910", Age: 28},
		DemoUser{Name: "邹许洌", Phone: "18345654911", Age: 28},
		DemoUser{Name: "梁晶茹", Phone: "18345654912", Age: 28},
		DemoUser{Name: "刘语林", Phone: "18345654913", Age: 28},
		DemoUser{Name: "曹可可", Phone: "18345654914", Age: 28},
	}
	many, _ := collection.InsertMany(
		context.TODO(),
		users,
	)
	fmt.Println(many)
}

// GetAll 查询全部数据
func GetAll(client *mongo.Client) {
	collection := client.Database(Database).Collection(Collection)
	var result []DemoUser
	cursor, _ := collection.Find(context.TODO(), bson.D{})
	if cursor.All(context.TODO(), &result) != nil {
		fmt.Println("数据查询错误")
	}
	for i, user := range result {
		UserFmt(i+1, user)
	}
}

// GetOne 根据条件查询返回一个结果
func GetOne() {
	collection := clientTest.Database(Database).Collection(Collection)
	res := DemoUser{}
	filter := bson.D{{"phone", "18345654908"}}
	err := collection.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		fmt.Println(err)
	}
	UserFmt(1, res)
}

// UserFmt 格式化输出
func UserFmt(index int, user DemoUser) {
	fmt.Printf("[%d] --姓名: %s -- 手机号: %s --- 年龄: %d--\n", index, user.Name, user.Phone, user.Age)
}

// UpdateById 根据ID更新数据
func UpdateById(client *mongo.Client) {
	collection := client.Database(Database).Collection(Collection)
	id, _ := primitive.ObjectIDFromHex("62a439b12f8ea2d0dc320146")
	filter := bson.D{{"_id", id}}
	user := DemoUser{}
	collection.FindOne(context.TODO(), filter).Decode(&user)
	fmt.Print("更新前：")
	UserFmt(1, user)
	updateRes, err := collection.UpdateByID(context.TODO(), id, bson.D{{"$set", bson.D{{"name", "张三(改)"}}}})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(updateRes)

	collection.FindOne(context.TODO(), filter).Decode(&user)
	fmt.Print("更新后：")
	UserFmt(1, user)
}

// UpdateOne 根据手机号更新数据
func UpdateOne() {
	collection := clientTest.Database(Database).Collection(Collection)
	filter := bson.D{{"phone", "13659630001"}}
	user := DemoUser{}
	collection.FindOne(context.TODO(), filter).Decode(&user)
	fmt.Print("更新前：")
	UserFmt(1, user)
	updateRes, err := collection.UpdateOne(context.TODO(), filter, bson.D{{"$set", bson.D{{"name", "汪鸣晨"}}}})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(updateRes)

	collection.FindOne(context.TODO(), filter).Decode(&user)
	fmt.Print("更新后：")
	UserFmt(1, user)
}

func UpdateMany(client *mongo.Client) {
	collection := client.Database(Database).Collection(Collection)
	filter := bson.D{{"phone", bson.D{{"$regex", "^1834565490"}}}}
	var users []DemoUser
	find, err := collection.Find(context.TODO(), filter)
	find.All(context.TODO(), &users)
	fmt.Print("更新前：")
	for i, user := range users {
		UserFmt(i+1, user)
	}
	updateRes, err := collection.UpdateMany(context.TODO(), filter, bson.D{{"$set", bson.D{{"age", 30}}}})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(updateRes)

	find, err = collection.Find(context.TODO(), filter)
	find.All(context.TODO(), &users)
	fmt.Print("更新后：")
	for i, user := range users {
		UserFmt(i+1, user)
	}
}

func ReplaceOne(client *mongo.Client) {
	collection := client.Database(Database).Collection(Collection)
	filter := bson.D{{"phone", "18346566786"}}
	one, err := collection.ReplaceOne(context.TODO(), filter, DemoUser{Name: "ReplaceOne"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(one)
}

// DeleteOne 根据手机号删除数据
func DeleteOne() {
	collection := clientTest.Database(Database).Collection(Collection)
	filter := bson.D{{"phone", "13659630004"}}
	one, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(one)
}

func Like(client *mongo.Client) {
	collection := client.Database(Database).Collection(Collection)
	var name = "雪"
	filter := bson.D{{"name", bson.D{{"$regex", "^.*" + name + ".*$"}}}}
	var users []DemoUser
	find, _ := collection.Find(context.TODO(), filter)
	find.All(context.TODO(), &users)
	for i, user := range users {
		UserFmt(i+1, user)
	}
}

func Page(client *mongo.Client, limit, page int64) {
	collection := client.Database(Database).Collection(Collection)
	var findOptions = &options.FindOptions{}
	findOptions.SetLimit(limit)
	findOptions.SetSkip(limit * page)
	findOptions.SetSort(bson.D{{"phone", 1}})
	cur, err := collection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		fmt.Println(err)
	}
	var users []DemoUser
	cur.All(context.TODO(), &users)
	for i, user := range users {
		UserFmt(i+1, user)
	}
}

func con() *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://julywhj:125846whj@cluster0.r1o1v.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	fmt.Println(client)
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	return client
}

// Ping 验证mongodb链接成功

func main() {
	client := con()
	//InstallMany(client)
	Page(client, 3, 1)
	Page(client, 3, 2)
	Page(client, 3, 3)
}
