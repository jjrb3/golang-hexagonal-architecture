package bootstrap

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/creating"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/platform/server"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/platform/storage/mysql"
)

const (
	host = "localhost"
	port = 3000

	dbUser = "root"
	dbPass = "root"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "codely"

	driverMySQL = "mysql"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open(driverMySQL, mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	creatingCourseService := creating.NewCourseService(courseRepository)

	srv := server.New(host, port, courseRepository, creatingCourseService)
	return srv.Run()
}
