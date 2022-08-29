package main

import (
	"log"
	"net"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	pb "github.com/octavio-luna/basic_grpc_api/proto"
	procedures "github.com/octavio-luna/basic_grpc_api/server/procedures"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting server...")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	serv := grpc.NewServer()

	s := procedures.NewServer()

	pb.RegisterECommerceServiceServer(serv, s)

	log.Println("Server started")
	if err = serv.Serve(listener); err != nil {
		panic("Can't serve, err: " + err.Error())
	}
}
