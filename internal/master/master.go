package master

import "Gungnir/internal/master/rest/server"

type Master struct {
	// clientPool
	// rpcclient
	httpServer *server.Server
}

func New() *Master {
	return &Master{}
}

func (m *Master) Start() error {

	// run http server
	err := m.httpServer.Run()

	if err != nil {
		return err
	}
	return nil
}
