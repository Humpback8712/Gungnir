package connector

import (
	G_mq "Gungnir/internal/master/G-mq"
	"log"
)

type Connector struct {
}

func New() *Connector {
	return &Connector{}
}
func (c *Connector) Connect(msg *G_mq.NodeConnectionMsg) {
	log.Println(msg)
}
