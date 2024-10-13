package main

import (
	"firewallBypass/shared"
	"github.com/hashicorp/yamux"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	log.Println("launching TPC server...")

	lis, err := net.Listen("tcp", ":3000")

	if err != nil {
		panic(err)
	}

	defer lis.Close()

	// ------------------
	log.Println("waiting TCP connections...")

	conn, err := lis.Accept()

	if err != nil {
		panic(err)
	}

	session, err := yamux.Client(conn, yamux.DefaultConfig())

	if err != nil {
		panic(err)
	}

	log.Println("launching gRPC server over TCP connection")

	server := grpc.NewServer()
	service := &MyCoreServer{}
	shared.RegisterCoreServerServer(server, service)

	go func() { err = server.Serve(session) }()
	yamuxNewConnections(lis)
}
