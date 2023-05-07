package server

import (
	"fmt"
	"net/http"

	"github.com/bypepe77/Rajoy-Hits-API/api/sounds"
	"github.com/gin-gonic/gin"
)

type server struct {
	config *Config
	engine *gin.Engine
}

func NewServer(config *Config) *server {
	return &server{
		config: config,
		engine: gin.New(),
	}
}

func (s *server) buildConnectionString() string {
	return fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
}

func (s *server) Run() error {
	s.registerRoutes()
	return s.engine.Run()
}

func (s *server) registerRoutes() {
	s.engine.GET("/", healthCheck)
	soundsRoutes := sounds.NewSoundRoutes(*s.engine.Group("/sounds"))
	soundsRoutes.RegisterSoundsRoutes()
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "pong"})
}
