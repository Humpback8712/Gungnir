package grpcserver

import (
	"Gungnir/pkg/grpc/pb_go"
	"google.golang.org/grpc"
	"net"
)

func RunGRPC() error {
	grpcServer := grpc.NewServer()
	AgentServer := &AgentServer{}

	pb_go.RegisterHealthServer(grpcServer, AgentServer)

	listen, err := net.Listen("tcp", "8889")
	if err != nil {
		return err
	}

	return grpcServer.Serve(listen)
}
