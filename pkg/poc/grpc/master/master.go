package main

import (
	"Gungnir/pkg/grpc/pb_go"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8889", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb_go.NewHealthClient(conn)

	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stream, err := c.Keep(ctx)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		for {
			err = stream.Send(&pb_go.HealthCheckRequest{ClusterName: "hi"})
			if err != nil {
				fmt.Println(err)
			}
			recv, err2 := stream.Recv()
			if err2 != nil {
				log.Println(err2)
				return
			}
			log.Println(recv)
		}
	}()
	select {}
}
