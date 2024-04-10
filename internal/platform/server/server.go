package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	mooc "github.com/gmd3v/gohex/internal"
	"github.com/gmd3v/gohex/internal/platform/server/handler/courses"
	"github.com/gmd3v/gohex/internal/platform/server/handler/health"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// deps
	courseRepository mooc.CourseRepository
}

func New(host string, port uint) *Server {
	return &Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.Default(),
	}
}

func (s *Server) Run() error {
	return s.engine.Run(s.httpAddr)
}

func (s *Server) RegisterRoutes() {
	s.engine.GET("/health", health.CheckHealth())
	s.engine.POST("/courses", courses.CreateHandler(s.courseRepository))
}

func (s *Server) Handle(method, path string, handler gin.HandlerFunc) {
	s.engine.Handle(method, path, handler)
}
