package agent

import (
	"Gungnir/pkg/grpc/pb_go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

func Connect(c *gin.Context) {
	cc, err := grpc.Dial("", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.String(http.StatusBadRequest, "err")
	}
	_ = pb_go.NewHealthClient(cc)

	c.String(http.StatusOK, "ok")
}
