package clientpool

import (
	"Gungnir/pkg/grpc/pb_go"
	"google.golang.org/grpc"
	"sync"
)

var CP *ClientPool

type GrpcClient struct {
	HealthClient pb_go.HealthClient
}
type ClientStatus struct {
	Cpu    float64
	Memory float64
}

type ClientPool struct {
	mu            *sync.RWMutex
	clients       map[string]*GrpcClient
	clientsStatus map[string]*ClientStatus
}

func New() *ClientPool {
	return &ClientPool{
		mu:            new(sync.RWMutex),
		clients:       make(map[string]*GrpcClient),
		clientsStatus: make(map[string]*ClientStatus),
	}
}

func (cp *ClientPool) AddClient(addr string, conn *grpc.ClientConn) {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	client := &GrpcClient{
		HealthClient: pb_go.NewHealthClient(conn),
	}
	cp.clients[addr] = client
}

func (cp *ClientPool) RemoveClient(addr string) {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	cp.clients[addr] = nil
}

func (cp *ClientPool) GetClient(addr string) (*GrpcClient, bool) {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	client, ok := cp.clients[addr]
	return client, ok
}

func (cp *ClientPool) UpdateClientStatus(addr string, status *ClientStatus) {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	cp.clientsStatus[addr] = status
}

func (cp *ClientPool) GetClientStatus(addr string) (*ClientStatus, bool) {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	client, ok := cp.clientsStatus[addr]
	return client, ok
}

func GetClientPoolOr() *ClientPool {
	if CP == nil {
		CP = New()
	}
	return CP
}
