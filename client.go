package main

import (
	"context"
	"log"
	"testing/grpc-demo/chat"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial: %v", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from Golang client",
	}

	res, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("SayHello: %v", err)
	}

	log.Printf("Response from server: %v", res.Body)
}
