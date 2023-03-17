package G_mq

var GMQ *Gmq

type NodeConnectionMsg struct {
	Name string
	Addr string
}

type Gmq struct {
	NodeConnectionChan chan *NodeConnectionMsg
}

func New() {
	GMQ = &Gmq{
		NodeConnectionChan: make(chan *NodeConnectionMsg),
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
