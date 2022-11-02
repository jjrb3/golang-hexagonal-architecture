package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	mooc "github.com/jjrb3/golang-hexagonal-architecture/internal"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/increasing"
	"github.com/kelseyhightower/envconfig"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/creating"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/platform/bus/inmemory"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/platform/server"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/platform/storage/mysql"
)

type config struct {
	// Server configuration.
	Host            string        `default:"localhost"`
	Port            uint          `default:"3000"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration.
	DBUser    string        `default:"root"`
	DBPass    string        `default:"root"`
	DBHost    string        `default:"localhost"`
	DBPort    uint          `default:"3306"`
	DBName    string        `default:"codely"`
	DBTimeout time.Duration `default:"5s"`
	// Drivers.
	DriverMySQL string `default:"mysql"`
}

func Run() error {
	var cfg config
	err := envconfig.Process("MOOC", &cfg)
	if err != nil {
		return err
	}

	mysqlURI := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	db, err := sql.Open(cfg.DriverMySQL, mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)

	courseRepository := mysql.NewCourseRepository(db, cfg.DBTimeout)

	creatingCourseService := creating.NewCourseService(courseRepository, eventBus)
	increasingCourseCounterService := increasing.NewCourseCounterService()

	createCourseCommandHandler := creating.NewCourseCommanderHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	eventBus.Subscribe(
		mooc.CourseCreatedEventType,
		creating.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseCounterService),
	)

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, courseRepository, commandBus)
	return srv.Run(ctx)
}
