package main

import (
	"context"
	"log"
	"testing/grpc-demo/chat"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("zoscan2b.pok.stglabs.ibm.com:9321", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial: %v", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.TextMessage{
		Body: "Hello from Golang client",
	}

	res, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("SayHello: %v", err)
	}

	log.Printf("Response from server: %v", res.Body)
}
