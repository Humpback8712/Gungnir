package proto

import (
	"Gungnir/pkg/grpc/pb_go"
	"io"
	"log"
)

type AgentServer struct {
	pb_go.UnimplementedHealthServer
}

func (s *AgentServer) Keep(stream pb_go.Health_KeepServer) error {
	for {
		// 接收流式请求
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Println(in)
		// 返回流式响应
		if err := stream.Send(&pb_go.HealthPackage{
			Info: "hello",
		}); err != nil {
			return err
		}
	}
}
