package connector

import (
	G_mq "Gungnir/internal/master/G-mq"
	"Gungnir/internal/master/nodemanager/clientpool"
	"Gungnir/internal/master/nodemanager/healthmanager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Connector struct {
}

func New() *Connector {
	return &Connector{}
}
func (c *Connector) Connect(msg *G_mq.NodeConnectionMsg) {

	conn, err := grpc.Dial(msg.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// add client into client pool
	pool := clientpool.GetClientPoolOr()
	pool.AddClient(msg.Addr, conn)

	go healthmanager.HeartBeatCheck(msg.Addr)
	select {}
}
