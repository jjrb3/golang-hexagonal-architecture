package mysql

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mooc "github.com/jjrb3/golang-hexagonal-architecture/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
)

const (
	courseID       = "37a0f027-15e6-47cc-a5d2-64183281087e"
	courseName     = "Test Course"
	courseDuration = "10 months"
)

func TestCourseRepository_Save_Error(t *testing.T) {
	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO courses (id, name, duration) VALUES (?, ?, ?)").
		WithArgs(courseID, courseName, courseDuration).
		WillReturnError(errors.New("something-failed"))

	repo := NewCourseRepository(db)

	err = repo.Save(context.Background(), course)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestCourseRepository_Save_Succeed(t *testing.T) {
	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO courses (id, name, duration) VALUES (?, ?, ?)").
		WithArgs(courseID, courseName, courseDuration).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewCourseRepository(db)

	err = repo.Save(context.Background(), course)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}
