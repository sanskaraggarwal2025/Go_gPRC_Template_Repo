package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"first_service/api"
	"first_service/dao"
	"first_service/proto"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sqlx.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	if err := dao.Migrate(db); err != nil {
		log.Fatalf("failed to migrate db: %v", err)
	}

	sqliteDAO := dao.NewSQLiteDAO(db)

	//create a new grpc server
	server := grpc.NewServer()

	proto.RegisterServiceServer(server, &api.ServiceServer{
		Dao: sqliteDAO,
	})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server.Serve(lis)
}
