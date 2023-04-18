package grpcserver

import (
	"Gungnir/pkg/grpc/pb_go"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func RunGRPC() error {
	grpcServer := grpc.NewServer()
	AgentServer := &AgentServer{}

	pb_go.RegisterHealthServer(grpcServer, AgentServer)
	go func() {
		time.Sleep(time.Second * 1)
		ConnectMaster()
	}()
	listen, err := net.Listen("tcp", ":8889")
	if err != nil {
		return err
	}

	return grpcServer.Serve(listen)
}

func ConnectMaster() {
	url := "http://127.0.0.1:8888/connect"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
