package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpcServer/server"

	"database/sql"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

var db2 *sql.DB

func connectWithConnector() (*sql.DB, error) {

	// Note: Saving credentials in environment variables is convenient, but not
	// secure - consider a more secure solution such as
	// Cloud Secret Manager (https://cloud.google.com/secret-manager) to help
	// keep passwords and other secrets safe.
	var (
		dbUser                 = " "                             // e.g. 'my-db-user'
		dbPwd                  = " "                       // e.g. 'my-db-password'
		dbName                 = " "                           // e.g. 'my-database'
		instanceConnectionName = " " // e.g. 'project:region:instance'
		usePrivate             = ""
	)

	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %w", err)
	}
	var opts []cloudsqlconn.DialOption
	if usePrivate != "" {
		opts = append(opts, cloudsqlconn.WithPrivateIP())
	}
	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return d.Dial(ctx, instanceConnectionName, opts...)
		})

	dbURI := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true",
		dbUser, dbPwd, dbName)

	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	fmt.Println("Connected to database")
	db2 = dbPool
	return dbPool, nil
}

type server struct {
	pb.UnimplementedGetInfoServer
}

const (
	port = ":3001"
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

	result, err := db2.Exec("INSERT INTO votos (name, album, year, ranked) VALUES (?, ?, ?, ?)", data.Name, data.Album, data.Year, data.Rank)
	if err != nil {
		// Manejar el error
		panic(err.Error())
	}

	// Obtener el número de filas afectadas por la inserción
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// Manejar el error
		panic(err.Error())
	}

	fmt.Printf("Se insertaron %d filas correctamente.\n", rowsAffected)

	// insertRedis(data)
	return &pb.ReplyInfo{Info: "--"}, nil
}

func printttt() {
	fmt.Println("->", db2)
}

func main() {
	db, err := connectWithConnector()
	if err != nil {
		log.Fatalf("connectWithConnector: %v", err)
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	printttt()

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	// redisConnect()

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
