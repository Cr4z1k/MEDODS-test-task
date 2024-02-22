package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, ginEngine *gin.Engine) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        ginEngine,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
	}

	return s.httpServer.ListenAndServe()
}
