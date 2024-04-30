package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpcServer/server"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

const (
	port   = ":3001"
	broker = "kafka:9092"
	topic  = "votes"
)

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	fmt.Println("-voto para: ", in.GetName())
	data := Data{
		Name:  in.GetName(),
		Year:  in.GetYear(),
		Album: in.GetAlbum(),
		Rank:  in.GetRank(),
	}

	// Kafka
	conn, err := kafka.DialLeader(context.Background(), "tcp", broker, topic, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
		fmt.Println("-err")
	}
	defer conn.Close()

	// Produce el mensaje
	message := fmt.Sprintf("%s,%s,%s,%s", data.Name, data.Album, data.Year, data.Rank)
	_, err = conn.WriteMessages(
		kafka.Message{
			Value: []byte(message),
		},
	)
	if err != nil {
		log.Fatalf("Error producing message to Kafka: %v", err)
		fmt.Println("-err")
	}

	fmt.Println("Message sent to Kafka!")

	return &pb.ReplyInfo{Info: "--"}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Server running on port", port)

	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
