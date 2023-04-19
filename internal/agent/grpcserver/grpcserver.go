package grpcserver

import (
	"Gungnir/pkg/grpc/pb_go"
)

type AgentServer struct {
	pb_go.UnimplementedHealthServer
}
