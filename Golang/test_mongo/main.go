package main

import (
	"fmt"
	"time"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func main() {
	//Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		fmt.Printf("client.Connect failed! error:%v", err.Error())
		return
	}

	//Test ping
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		fmt.Printf("client.Ping failed! error:%v", err.Error())
		return
	}

	//Get collection
	collection := client.Database("test").Collection("test")

	//Get all
	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		fmt.Printf("collection.Find failed! error:%v", err.Error())
		return
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)

		if err != nil {
			fmt.Printf("cur.Decode failed! error:%v", err.Error())
			continue
		}

		// do something with result....
		fmt.Printf("result:%v", result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	//Update
	_, err = collection.UpdateOne(ctx, bson.M{"a":99.99}, bson.M{"b":"test b1"}, )

	if err != nil {
		fmt.Printf("collection.UpdateOne failed! error:%v", err.Error())
		return
	}
}
