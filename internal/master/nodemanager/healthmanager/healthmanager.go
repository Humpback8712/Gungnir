package healthmanager

import (
	"Gungnir/internal/master/nodemanager/clientpool"
	"Gungnir/pkg/grpc/pb_go"
	"context"
	"log"
	"time"
)

type HealthManager struct {
}

func New() *HealthManager {
	return &HealthManager{}
}

func HeartBeatCheck(addr string) {

	// get client
	clientPool := clientpool.GetClientPoolOr()
	client, _ := clientPool.GetClient(addr)

	for {
		stream, err := client.HealthClient.Keep(context.Background())
		if err != nil {
			log.Println("err:", err)
			return
		}
		err = stream.Send(&pb_go.HealthCheckRequest{
			ClusterName: "master",
		})
		if err != nil {
			log.Println(err)
		}
		recv, err2 := stream.Recv()
		if err2 != nil {
			log.Println(err2)
		}
		clientPool.UpdateClientStatus(addr, &clientpool.ClientStatus{
			Cpu:    float64(recv.Resource.CpuRest),
			Memory: float64(recv.Resource.DiskRest),
		})
		time.Sleep(time.Second * 3)
	}
}
