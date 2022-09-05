package mysql

const (
	sqlCourseTable = "courses"
)

// sqlCourse is a data mapper.
type sqlCourse struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Duration string `db:"duration"`
}
