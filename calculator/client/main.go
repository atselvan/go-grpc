package main

import (
	"log"

	"github.com/atselvan/go-grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:5002"

func main() {
	client, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to initialise client: %v\n", err)
	}
	defer client.Close()

	c := proto.NewCalculatorServiceClient(client)
	add(c, 5, 5)
	sub(c, 12, 2)
	mul(c, 5, 2)
	div(c, 20, 2)
}
