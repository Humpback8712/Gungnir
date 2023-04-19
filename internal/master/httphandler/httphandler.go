package httphandler

import (
	"Gungnir/internal/master/httphandler/nodehandler"
	"Gungnir/internal/master/httpserver"
	"github.com/gin-gonic/gin"
)

var HttpHandler *Httphandler

type Httphandler struct {
	e *gin.Engine
	// handlers
	nodeHandler *nodehandler.Nodehandler
}

func New() {
	HttpHandler = &Httphandler{
		e:           httpserver.GetHttpServerOr().GetEngine(),
		nodeHandler: nodehandler.New(),
	}
	HttpHandler.Register()
}

func GetHttpHandlerOr() *Httphandler {
	if HttpHandler == nil {
		New()
	}
	return HttpHandler
}

func (h *Httphandler) Register() {
	
	h.e.GET("hi", h.nodeHandler.Hi)
	h.e.GET("connect", h.nodeHandler.Connect)
	h.e.POST("upload/job", h.nodeHandler.UploadJob)

}
