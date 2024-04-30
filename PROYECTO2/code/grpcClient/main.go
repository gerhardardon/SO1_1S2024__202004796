package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpcClient/client"
)

var ctx = context.Background()

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func insertData(c *fiber.Ctx) error {
	var data map[string]string
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	ranks := Data{
		Name:  data["name"],
		Album: data["album"],
		Year:  data["year"],
		Rank:  data["rank"],
	}

	conn, err := grpc.Dial("grpc-server:3001", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
		fmt.Println("-err")
	}
	fmt.Println("Conectado al server")

	cl := pb.NewGetInfoClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
			fmt.Println("-err")
		}
	}(conn)

	ret, err := cl.ReturnInfo(ctx, &pb.RequestId{
		Name:  ranks.Name,
		Album: ranks.Album,
		Year:  ranks.Year,
		Rank:  ranks.Rank,
	})
	if err != nil {
		log.Fatalln(err)
		fmt.Println("-err")
	}

	fmt.Println("Respuesta del server " + ret.GetInfo())

	return nil
}

/*func convertToInt(s string) {
	panic("unimplemented")
}*/

func main() {
	app := fiber.New()
	app.Post("/insert", insertData)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	err := app.Listen(":3000")
	if err != nil {
		log.Fatalln(err)
		fmt.Println("Error")
		return
	}
}
