package main

import "Gungnir/internal/agent/grpcserver"

func main() {
	err := grpcserver.RunGRPC()
	if err != nil {
		panic(err)
	}
}
