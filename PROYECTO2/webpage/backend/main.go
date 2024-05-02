package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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
	app := fiber.New()
	ctx := context.Background()

	//mongo connection
	clientOptions := options.Client().ApplyURI("mongodb://35.225.122.166:27017") // Change this IP to the external IP of the mongoDB instance
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
	fmt.Println("Connected to MongoDB")
	defer mongoClient.Disconnect(ctx)

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/records", func(c *fiber.Ctx) error {
		collection := mongoClient.Database("votesdb").Collection("votes")
		findOptions := options.Find()
		findOptions.SetLimit(20)
		findOptions.SetSort(bson.D{{Key: "_id", Value: -1}})

		cur, err := collection.Find(ctx, bson.D{}, findOptions)
		if err != nil {
			log.Fatal("Error finding records", err)
			return err
		}

		var records []Data
		if err := cur.All(ctx, &records); err != nil {
			log.Fatal("Error decoding records", err)
			return err
		}

		return c.JSON(records)
	})

	app.Listen(":8080")
}
