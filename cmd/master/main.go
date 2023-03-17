package main

import (
	"Gungnir/internal/master"
	"github.com/sirupsen/logrus"
)

func main() {
	//cc, err := grpc.Dial("http://127.0.0.1:8889", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	panic(err)
	//}
	//log.Println("agent" + "--" + "dial success")
	//hc := pb_go.NewHealthClient(cc)
	//keep, err := hc.Keep(context.Background())
	//log.Println(keep)
	//if err != nil {
	//	panic(err)
	//}
	//return

	err := master.New().Start()
	if err != nil {
		logrus.Fatal(err)
	}
}
