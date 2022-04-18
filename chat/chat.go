package chat

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *TextMessage) (*TextMessage, error) {
	log.Printf("Received Message from Client: %s", message.Body)
	return &TextMessage{Body: "Hello from Server!"}, nil
}
