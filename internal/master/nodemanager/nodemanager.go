package nodemanager

import (
	G_mq "Gungnir/internal/master/G-mq"
	"Gungnir/internal/master/nodemanager/clientpool"
	"Gungnir/internal/master/nodemanager/connector"
	"Gungnir/internal/master/nodemanager/healthmanager"
)

var NodeManager *Nodemanager

type Nodemanager struct {
	ClientPool    *clientpool.ClientPool
	HealthManager *healthmanager.HealthManager
	Connector     *connector.Connector
}

func New() {
	NodeManager = &Nodemanager{
		ClientPool:    clientpool.New(),
		HealthManager: healthmanager.New(),
		Connector:     connector.New(),
	}
}

func GetNodeMangerOr() *Nodemanager {
	if NodeManager == nil {
		New()
	}
	return NodeManager
}

func (n *Nodemanager) Run() {
	go func() {
		for {
			msg := <-G_mq.GetMqOr().GetNodeConnectChan()
			go n.Connector.Connect(msg)
		}
	}()
}
