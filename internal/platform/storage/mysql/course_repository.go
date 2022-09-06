package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/huandu/go-sqlbuilder"
	mooc "github.com/jjrb3/golang-hexagonal-architecture/internal"
)

// CourseRepository is a MySQL mooc.CourseRepository implementation.
type CourseRepository struct {
	db *sql.DB
}

// NewCourseRepository initialize a MySQL-based implementation of mooc.CourseRepository.
func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
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

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}

func (r *CourseRepository) FindAll(ctx context.Context) ([]mooc.Course, error) {
	sb := sqlbuilder.NewStruct(new(sqlCourse))

	query, args := sb.SelectFrom(sqlCourseTable).Build()

	records, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return []mooc.Course{}, err
	}

	log.Println(records)

	return []mooc.Course{}, nil
}
