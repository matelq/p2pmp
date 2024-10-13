package main

import (
	"context"
	"github.com/hashicorp/yamux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

var connectionDict = map[string]*grpc.ClientConn{}
var connectionChanDict = map[string]chan *grpc.ClientConn{}

func yamuxNewConnections(lis net.Listener) {
	for clientIndex := 0; ; clientIndex++ {
		log.Println("waiting TCP connections for new clients...")

		conn, err := lis.Accept()

		if err != nil {
			panic(err)
		}

		if connectionDict[conn.RemoteAddr().String()] != nil {
			continue
		}

		session, err := yamux.Client(conn, yamux.DefaultConfig())

		if err != nil {
			panic(err)
		}

		clientConn, err := grpc.NewClient(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return session.Open() }))

		if err != nil {
			panic(err)
		}

		connectionDict[conn.RemoteAddr().String()] = clientConn
		connectionChanDict[conn.RemoteAddr().String()] = make(chan *grpc.ClientConn, 1)
		connectionChanDict[conn.RemoteAddr().String()] <- clientConn
	}
}
