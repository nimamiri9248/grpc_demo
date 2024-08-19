package main

import (
	"context"
	"io"
	"log"

	pb "github.com/nimamiri9248/grpc_sample/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	log.Printf("streaming has astarted")
	if err != nil {
		log.Fatalf("couldnot greet %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("couldnot greet %v", err)
		}
		log.Printf("%s", res.Message)
	}
	log.Printf("streaming has finished")
}
