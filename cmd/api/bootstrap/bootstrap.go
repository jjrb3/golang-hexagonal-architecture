package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/jjrb3/golang-hexagonal-architecture/internal/platform/server"
)

const (
	host = "localhost"
	port = 3000

	dbUser = "root"
	dbPass = "root"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "codely"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}
