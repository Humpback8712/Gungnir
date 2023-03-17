package httpserver

import "github.com/gin-gonic/gin"

var HttpServer *Httpserver

type Httpserver struct {
	Ip   string
	Port string
	e    *gin.Engine
}

func New() {
	// config msg
	HttpServer = &Httpserver{
		e:    gin.Default(),
		Port: ":8888",
	}
}

func GetHttpServerOr() *Httpserver {
	if HttpServer == nil {
		New()
	}
	return HttpServer
}

func (h *Httpserver) GetEngine() *gin.Engine {
	return h.e
}

func (h *Httpserver) Run() {
	// run server
	if err := h.e.Run(h.Port); err != nil {
		panic(err)
	}
}
