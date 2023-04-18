package connector

import (
	G_mq "Gungnir/internal/master/G-mq"
	"Gungnir/pkg/grpc/pb_go"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type Connector struct {
}

func New() *Connector {
	return &Connector{}
}
func (c *Connector) Connect(msg *G_mq.NodeConnectionMsg) {
	conn, err := grpc.Dial("127.0.0.1:8889", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb_go.NewHealthClient(conn)

	stream, err := client.Keep(context.Background())
	if err != nil {
		log.Println(err)
	}
	go func() {
		for {
			err = stream.Send(&pb_go.HealthCheckRequest{ClusterName: "localhost testing"})
			if err != nil {
				log.Println(err)
			}
			recv, err2 := stream.Recv()
			if err2 != nil {
				log.Println(err2)
			}
			log.Println(recv)
			time.Sleep(time.Second * 3)
		}
	}()
	select {}
}
