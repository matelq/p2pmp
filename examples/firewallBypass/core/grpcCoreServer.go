package main

import (
	"context"
	"firewallBypass/shared"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"log"
	"net"
)

var clientDict = map[string]*shared.ClientServerClient{}

func HandleRegistration(addr net.Addr) (clientChan chan shared.ClientServerClient)
{
	// Проверка на наличие этого клиента в мапе clientDict
	// достать из мап в yamuxLogic clientConn или chan clientConn
	// вернуть канал с созданным клиентом + записать его в clientDict
}

type MyCoreServer struct {
	shared.UnimplementedCoreServerServer
}

func (s MyCoreServer) RegisterMe(grpcContext context.Context, message *shared.RegisterRequest) (*shared.RegisterResponse, error) {
	p, _ := peer.FromContext(grpcContext)
	log.Printf("Got RegisterMe request from: %v", p.Addr.String())
	p.Addr.String()
}

func (s MyCoreServer) SendMessageCore(grpcContext context.Context, message *shared.MessageCore) (*shared.EchoCore, error) {
	log.Printf("SendMessageCore called: %s", message.Text)

	// TODO: валидацию айдишника
	//client := clientDict[message.Id]

	res, err := client.SendMessageToClient(grpcContext, &shared.MessageClient{Text: "Message from server to client"})

	if err != nil {
		panic(err)
	}

	log.Printf("From client to server: %s", res.Text)

	return &shared.EchoCore{
		Text: fmt.Sprintf("Echo: %s", message.Text),
	}, nil
}
