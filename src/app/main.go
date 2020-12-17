package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	pb "test-chat/src/app/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()

	chatServer := NewChatServiceServer()
	pb.RegisterChatServiceServer(server, chatServer)

	reflection.Register(server)

	log.Print("Server Ready")

	if err := server.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
