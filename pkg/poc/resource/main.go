package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"time"
)

func main() {
	usage, err := disk.Usage("/")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Total: %.2f Megabyte\n", float64(usage.Total)/(1024*1024))
	fmt.Printf("Used: %.2f Megabyte\n", float64(usage.Used)/(1024*1024))
	fmt.Printf("Free: %.2f Megabyte\n", float64(usage.Free)/(1024*1024))

	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("CPU percent: %.2f%%\n", cpuPercent[0])

	numCPU, _ := cpu.Counts(false) // 获取总共的 CPU 数量

	cpuUsage, _ := cpu.Percent(time.Second, false) // 获取当前 CPU 使用率

	usedCPU := float64(numCPU) * (cpuUsage[0] / 100) // 计算已使用的 CPU 核心数量

	availableCPU := float64(numCPU) - usedCPU // 计算剩余的 CPU 核心数量

	fmt.Printf("Total CPU: %d\n", numCPU)
	fmt.Printf("Available CPU: %.2f\n", availableCPU)
}
