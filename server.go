package main

import (
	"log"
	"net"

	"testing/grpc-demo/chat"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	s := chat.Server{}
	chat.RegisterChatServiceServer(grpcServer, &s)

	log.Println("Listening on :9000...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Serve: %v", err)
	}
}
