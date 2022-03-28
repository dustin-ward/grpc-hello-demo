package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"testing/grpc-demo/chat"

	"google.golang.org/grpc"
)

func shutdown(s *grpc.Server) {
	var input string
	fmt.Scanln(&input)
	for !(input == "stop" || input == "STOP") {
		fmt.Scanln(&input)
	}
	s.GracefulStop()
	os.Exit(0)
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	s := chat.Server{}
	chat.RegisterChatServiceServer(grpcServer, &s)

	go shutdown(grpcServer)
	log.Println("Listening on :9000...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Serve: %v", err)
	}
}
