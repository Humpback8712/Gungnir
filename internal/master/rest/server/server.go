package server

import (
	"Gungnir/internal/master/rest/router"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	ip     string
	engine *gin.Engine
}

func New() *Server {
	return &Server{}
}

func (s *Server) Run() error {

	router.SetUpRouter(s.engine)

	if err := s.engine.Run(s.port); err != nil {
		return err
	}
	return nil
}
