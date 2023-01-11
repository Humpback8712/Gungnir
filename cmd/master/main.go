package main

import (
	"Gungnir/internal/master"
	"github.com/sirupsen/logrus"
)

func main() {
	m := master.New()
	err := m.Start()
	if err != nil {
		logrus.Fatal(err)
	}
}
