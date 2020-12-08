package mongoq

import (
	"context"
	"log"
	"os"
	"time"

	//一种数据格式
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

//init方法初始化日志存储文件和连接数据库
func init() {
	setLog()
}

func Conn(uri string, dbName string, colName string) {
	client, _ = mongo.NewClient(options.Client().ApplyURI(uri))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Println("[init]", err)
	}
	collection = client.Database(dbName).Collection(colName)
}

//FineOne封装了mongo的FindOne方法
func FineOne(cname string, database string, filter interface{}) bson.M {
	collection = client.Database(database).Collection(cname)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var result bson.M
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return bson.M{}
		}
		log.Println("[FindOne]", err)
	}
	return result
}

func InsertOne() {}

func DeleteOne() {}

func UpdateOne() {}

func setLog() {
	file := "./log.txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	log.SetPrefix("[db]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}
