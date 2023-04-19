package master

import (
	G_mq "Gungnir/internal/master/G-mq"
	"Gungnir/internal/master/httphandler"
	"Gungnir/internal/master/httpserver"
	"Gungnir/internal/master/nodemanager"
	"Gungnir/internal/master/nodemanager/clientpool"
)

type Master struct {
	httpServer  *httpserver.Httpserver
	httpHandler *httphandler.Httphandler
	nodeManager *nodemanager.Nodemanager
	clientpool  *clientpool.ClientPool
	gmq         *G_mq.Gmq
}

// New init all components
func New() *Master {
	return &Master{
		clientpool:  clientpool.New(),
		httpServer:  httpserver.GetHttpServerOr(),
		httpHandler: httphandler.GetHttpHandlerOr(),
		nodeManager: nodemanager.GetNodeMangerOr(),
		gmq:         G_mq.GetMqOr(),
	}
}

func (m *Master) Start() error {
	m.nodeManager.Run()
	m.httpServer.Run()
	return nil
}
