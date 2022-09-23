package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	mooc "github.com/jjrb3/golang-hexagonal-architecture/internal"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/platform/server/handler/courses"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/platform/server/handler/health"
	"github.com/jjrb3/golang-hexagonal-architecture/kit/command"
)

type Server struct {
	engine   *gin.Engine
	httpAddr string

	// deps.
	commandBus       command.Bus
	courseRepository mooc.CourseRepository
}

func New(
	host string,
	port uint,
	courseRepository mooc.CourseRepository,
	commandBus command.Bus,
) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		commandBus:       commandBus,
		courseRepository: courseRepository,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/course", courses.CreateHandler(s.commandBus))
	s.engine.GET("/courses", courses.FindAllHandler(s.courseRepository))
}
