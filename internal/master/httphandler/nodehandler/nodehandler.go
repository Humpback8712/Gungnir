package nodehandler

import (
	G_mq "Gungnir/internal/master/G-mq"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Nodehandler struct {
}

func New() *Nodehandler {
	return &Nodehandler{}
}

func (n *Nodehandler) Hi(c *gin.Context) {
	//G_mq.GetMqOr().NodeConnectionChan <- &G_mq.NodeConnectionMsg{}
	c.String(http.StatusOK, "ok")
}

func (n *Nodehandler) Connect(c *gin.Context) {
	G_mq.GetMqOr().NodeConnectionChan <- &G_mq.NodeConnectionMsg{
		Name: "12",
		Addr: "127.0.0.1:8899",
	}
	c.String(http.StatusOK, "plz wait")
}
