package G_mq

var GMQ *Gmq

type NodeConnectionMsg struct {
	Name string
	Addr string
}
type JobStartSchedulerMsg struct {
	Name   string
	Path   string
	Cpu    int
	Memory int
}

type Gmq struct {
	NodeConnectionChan    chan *NodeConnectionMsg
	JobStartSchedulerChan chan *JobStartSchedulerMsg
}

func New() {
	GMQ = &Gmq{
		NodeConnectionChan:    make(chan *NodeConnectionMsg),
		JobStartSchedulerChan: make(chan *JobStartSchedulerMsg),
	}
}

func GetMqOr() *Gmq {
	if GMQ == nil {
		New()
	}
	return GMQ
}

func (g *Gmq) GetNodeConnectChan() chan *NodeConnectionMsg {
	return g.NodeConnectionChan
}
