package main

import (
	"log"

	pb "github.com/atselvan/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:5001"

func main() {
	client, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to initialise client: %v\n", err)
	}
	defer client.Close()

	c := pb.NewGreetServiceClient(client)
	doGreet(c)
}
