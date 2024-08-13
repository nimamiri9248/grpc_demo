package main

import (
	"log"

	pb "github.com/nimamiri9248/grpc_sample/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't connect %v", err)

	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// names := &pb.NamesList
	callsayHello(client)

}
