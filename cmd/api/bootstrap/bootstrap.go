package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	mooc "github.com/jjrb3/golang-hexagonal-architecture/internal"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/increasing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/creating"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/platform/bus/inmemory"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/platform/server"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/platform/storage/mysql"
)

const (
	host = "localhost"
	port = 3000

	dbUser    = "root"
	dbPass    = "root"
	dbHost    = "localhost"
	dbPort    = "3306"
	dbName    = "codely"
	dbTimeout = 5 * time.Second

	driverMySQL = "mysql"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open(driverMySQL, mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)

	courseRepository := mysql.NewCourseRepository(db, dbTimeout)

	creatingCourseService := creating.NewCourseService(courseRepository, eventBus)
	increasingCourseCounterService := increasing.NewCourseCounterService()

	createCourseCommandHandler := creating.NewCourseCommanderHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	eventBus.Subscribe(
		mooc.CourseCreatedEventType,
		creating.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseCounterService),
	)

	ctx, srv := server.New(context.Background(), host, port, courseRepository, commandBus)
	return srv.Run(ctx)
}
