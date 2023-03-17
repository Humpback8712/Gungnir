package main

import (
	"Gungnir/internal/agent/grpcserver"
	"Gungnir/pkg/grpc/pb_go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {

	grpcServer := grpc.NewServer()
	AgentServer := &grpcserver.AgentServer{}

	pb_go.RegisterHealthServer(grpcServer, AgentServer)

	listen, err := net.Listen("tcp", ":8889")
	if err != nil {
		panic(err)
	}
	err = grpcServer.Serve(listen)
	if err != nil {
		logrus.Fatal(err)
	}
	return
}
