package grpcserver

import (
	"Gungnir/pkg/grpc/pb_go"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"io"
	"log"
	"time"
)

func (s *AgentServer) Keep(stream pb_go.Health_KeepServer) error {
	for {
		// 接收流式请求
		in, err := stream.Recv()
		if err == io.EOF {
			log.Println(err)
			return nil
		}
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println(in)
		// 返回流式响应
		if err := stream.Send(&pb_go.HealthPackage{
			ClusterName: "test cluster",
			Status:      "1",
			Info:        "Everything is doing ok",
			Resource:    s.CheckResource(),
		}); err != nil {
			return err
		}
	}
}

func (s *AgentServer) CheckResource() *pb_go.Resource {
	usage, err := disk.Usage("/")
	if err != nil {
		fmt.Println(err)
	}
	diskTotal := float64(usage.Total) / (1024 * 1024)
	diskRest := float64(usage.Free) / (1024 * 1024)

	numCPU, _ := cpu.Counts(false) // 获取总共的 CPU 数量

	cpuUsage, _ := cpu.Percent(time.Second, false) // 获取当前 CPU 使用率

	usedCPU := float64(numCPU) * (cpuUsage[0] / 100) // 计算已使用的 CPU 核心数量

	availableCPU := float64(numCPU) - usedCPU // 计算剩余的 CPU 核心数量

	fmt.Printf("Total CPU: %d\n", numCPU)
	fmt.Printf("Available CPU: %.2f\n", availableCPU)
	return &pb_go.Resource{
		CpuTotal:  float32(numCPU),
		CpuRest:   float32(availableCPU),
		DiskTotal: float32(diskTotal),
		DiskRest:  float32(diskRest),
	}
}
