package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	mooc "github.com/jjrb3/golang-hexagonal-architecture/internal"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/platform/server/handler/courses"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/platform/server/handler/health"
	"github.com/jjrb3/golang-hexagonal-architecture/kit/command"
)

type Server struct {
	engine   *gin.Engine
	httpAddr string

	shutdownTimeout time.Duration

	// deps.
	commandBus       command.Bus
	courseRepository mooc.CourseRepository
}

func New(
	ctx context.Context,
	host string,
	port uint,
	courseRepository mooc.CourseRepository,
	commandBus command.Bus,
) (context.Context, Server) {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		commandBus:       commandBus,
		courseRepository: courseRepository,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/course", courses.CreateHandler(s.commandBus))
	s.engine.GET("/courses", courses.FindAllHandler(s.courseRepository))
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout) // 10 seconds.
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		<-c
		cancel()
	}()

	return ctx
}
