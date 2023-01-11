package router

import (
	"Gungnir/internal/master/rest/handler/agent"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(e *gin.Engine) {
	e.GET("connect", agent.Connect)
}
