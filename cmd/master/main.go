package main

import (
	"Gungnir/internal/master"
	"github.com/sirupsen/logrus"
)

func main() {
	err := master.New().Start()
	if err != nil {
		logrus.Fatal(err)
	}
}
