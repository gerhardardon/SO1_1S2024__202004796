package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MusicRecord struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func main() {
	ctx := context.Background()

	fmt.Println("Conectando a Kafka...")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"kafka:9092"},
		Topic:     "votes",
		GroupID:   "grupo-consumidor",
		Partition: 0,
	})
	defer r.Close()

	fmt.Println("Conectando a Redis...")
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379", // host:port del servicio Redis
	})

	fmt.Println("Conectando a MongoDB...")
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		fmt.Printf("failed to connect to mongo: %v\n", err)
		return
	}
	mongoCollection := mongoClient.Database("testdb").Collection("records")

	fmt.Println("Iniciando la recepción de mensajes desde Kafka...")
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Printf("failed to read message: %v\n", err)
			continue
		}
		fmt.Printf("Mensaje recibido en Kafka con offset %d: clave=%s, valor=%s\n", m.Offset, string(m.Key), string(m.Value))

		parts := strings.Split(string(m.Value), ",")
		if len(parts) < 4 {
			fmt.Println("Mensaje recibido no tiene el formato esperado: se esperaban al menos 4 partes")
			continue
		}

		record := MusicRecord{
			Name:  strings.TrimSpace(parts[0]),
			Album: strings.TrimSpace(parts[1]),
			Year:  strings.TrimSpace(parts[2]),
			Rank:  strings.TrimSpace(parts[3]),
		}

		fmt.Println("Datos deserializados del mensaje:")
		fmt.Printf("Nombre: %s\n", record.Name)
		fmt.Printf("Álbum: %s\n", record.Album)
		fmt.Printf("Año: %s\n", record.Year)
		fmt.Printf("Rango: %s\n", record.Rank)

		// Incrementa un contador de votos para la banda en Redis bajo el hash 'teams'
		teamKey := record.Name
		_, err = rdb.HIncrBy(ctx, "bands", teamKey, 1).Result() //llave para grafana hgetall:redis
		if err != nil {
			fmt.Printf("failed to increment vote count for %s: %v\n", teamKey, err)
			continue
		}

		// Incrementa el contador total de votos
		_, err = rdb.Incr(ctx, "total:votes").Result()
		if err != nil {
			fmt.Printf("failed to increment total votes: %v\n", err)
			continue
		}

		fmt.Printf("Votos incrementados para '%s' y total general actualizado.\n", record.Name)

		// Guarda los datos en MongoDB
		_, err = mongoCollection.InsertOne(ctx, bson.M{
			"name":  record.Name,
			"album": record.Album,
			"year":  record.Year,
			"rank":  record.Rank,
		})
		if err != nil {
			fmt.Printf("failed to insert Mongo document: %v\n", err)
			continue
		}
		fmt.Printf("Datos guardados en MongoDB para '%s'.\n", record.Name)
	}
}
