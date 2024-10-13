package main

import (
	"context"
	"firewallBypass/shared"
	"fmt"
	"github.com/hashicorp/yamux"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type MyClientServer struct {
	shared.UnimplementedClientServerServer
}

func (s MyClientServer) SendMessageClient(context context.Context, message *shared.MessageClient) (*shared.EchoClient, error) {
	log.Printf("SendMessageClient called: %s", message.Text)

	return &shared.EchoClient{
		Text: fmt.Sprintf("Echo: %s", message.Text),
	}, nil
}

func main() {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:3000", 10*time.Second)

	if err != nil {
		panic(err)
	}

	session, err := yamux.Server(conn, yamux.DefaultConfig())

	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	service := &MyClientServer{}
	shared.RegisterClientServerServer(server, service)

	log.Println("launching gRPC server over TCP connection...")

	err = server.Serve(session)

	if err != nil {
		panic(err)
	}
}
