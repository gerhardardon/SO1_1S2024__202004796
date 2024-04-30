package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func main() {
	ctx := context.Background()

	//kafka conection
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"kafka:9092"},
		Topic:     "votes",
		Partition: 0,
		GroupID:   "my-group",
	})
	defer reader.Close()

	//redis connection
	redisdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Error connecting to redis", err)
		fmt.Println("Error connecting to redis")
	}
	defer redisdb.Close()

	//mongo connection
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error connecting to mongo", err)
		fmt.Println("Error connecting to mongo")
	}
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error connecting to mongo", err)
		fmt.Println("Error connecting to mongo")
	}
	defer mongoClient.Disconnect(ctx)
	mongoCollection := mongoClient.Database("votesdb").Collection("votes")

	//Read messages
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Fatal("Error reading message", err)
			fmt.Println("Error reading message")
			continue
		}
		fmt.Println("Message: ", string(msg.Value))

		data := strings.Split(string(msg.Value), ",")
		if len(data) != 4 {
			log.Fatal("Error in format", err)
			fmt.Println("Error in format")
			continue
		}

		vote := Data{
			Name:  strings.TrimSpace(data[0]),
			Album: strings.TrimSpace(data[1]),
			Year:  strings.TrimSpace(data[2]),
			Rank:  strings.TrimSpace(data[3]),
		}

		//Redis
		err = redisdb.HIncrBy(ctx, "bands", vote.Name, 1).Err()
		if err != nil {
			log.Fatal("Error incrementing redis", err)
			continue
		}

		//Mongo
		_, err = mongoCollection.InsertOne(ctx, vote)
		if err != nil {
			log.Fatal("Error inserting mongo", err)
			continue
		}
		log.Println("-vote saved ", vote)
	}
}
