package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nimamiri9248/grpc_sample/proto"
)

func callsayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("couldnot greet %v", err)

	}
	log.Printf("%s", resp.Message)
}
