package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/huandu/go-sqlbuilder"
	mooc "github.com/jjrb3/golang-hexagonal-architecture/internal"
)

// CourseRepository is a MySQL mooc.CourseRepository implementation.
type CourseRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewCourseRepository initialize a MySQL-based implementation of mooc.CourseRepository.
func NewCourseRepository(db *sql.DB, dbTimeout time.Duration) *CourseRepository {
	return &CourseRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Save implements the mooc.CourseRepository interface.
func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	sb := sqlbuilder.NewStruct(new(sqlCourse))

	query, args := sb.InsertInto(
		sqlCourseTable, sqlCourse{
			ID:       course.ID().String(),
			Name:     course.Name().String(),
			Duration: course.Duration().String(),
		},
	).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}

// FindAll return all courses.
func (r *CourseRepository) FindAll(ctx context.Context) ([]mooc.Course, error) {
	sb := sqlbuilder.NewStruct(new(sqlCourse))

	query, args := sb.SelectFrom(sqlCourseTable).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var courses []mooc.Course

	for rows.Next() {
		var sc sqlCourse
		var course mooc.Course

		err = rows.Scan(&sc.ID, &sc.Name, &sc.Duration)
		if err != nil {
			return nil, err
		}

		course, err = mooc.NewCourse(sc.ID, sc.Name, sc.Duration)
		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}
