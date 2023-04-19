package master

import (
	G_mq "Gungnir/internal/master/G-mq"
	G_scheduler "Gungnir/internal/master/G-scheduler"
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
	gscheduler  *G_scheduler.GScheduler
}

// New init all components
func New() *Master {
	return &Master{
		clientpool:  clientpool.GetClientPoolOr(),
		httpServer:  httpserver.GetHttpServerOr(),
		httpHandler: httphandler.GetHttpHandlerOr(),
		nodeManager: nodemanager.GetNodeMangerOr(),
		gmq:         G_mq.GetMqOr(),
		gscheduler:  G_scheduler.GetSchedulerOr(),
	}
}

func (m *Master) Start() error {
	go m.gscheduler.Run()
	m.nodeManager.Run()
	m.httpServer.Run()
	return nil
}
