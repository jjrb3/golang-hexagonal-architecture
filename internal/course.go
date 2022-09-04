package mooc

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

var ErrInvalidCourseID = errors.New("invalid Course ID")

// CourseID represents the course unique identifier.
type CourseID struct {
	value string
}

// NewCourseID instantiate the VO for CourseID.
func NewCourseID(value string) (CourseID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return CourseID{}, fmt.Errorf("%w: %s", ErrInvalidCourseID, value)
	}

	return CourseID{
		value: v.String(),
	}, nil
}

func (id CourseID) string() string {
	return id.value
}

var ErrEmptyCourseName = errors.New("the field Course Name can not be empty")
var ErrLengthCourseName = errors.New("the field Course Name would be greater than 10 characters")

// lengthCourseName is the minimum length to a course.
const lengthCourseName int = 10

// CourseName represents the course name.
type CourseName struct {
	value string
}

// NewCourseName instantiate VO for CourseName
func NewCourseName(value string) (CourseName, error) {
	if value == "" {
		return CourseName{}, ErrEmptyCourseName
	}

	if len(value) < lengthCourseName {
		return CourseName{}, ErrLengthCourseName
	}

	return CourseName{
		value: value,
	}, nil
}

// String type converts the CourseName into string.
func (name CourseName) String() string {
	return name.value
}

var ErrEmptyDuration = errors.New("the field Duration can not be empty")
var ErrDateDuration = errors.New("the date duration is incorrect")

// CourseDuration represents the course duration.
type CourseDuration struct {
	value string
}

func NewCourseDuration(value string) (CourseDuration, error) {
	if value == "" {
		return CourseDuration{}, ErrEmptyDuration
	}

	duration := strings.ToLower(value)

	if !(strings.Contains(duration, "months") ||
		strings.Contains(duration, "month") ||
		strings.Contains(duration, "day") ||
		strings.Contains(duration, "days") ||
		strings.Contains(duration, "year") ||
		strings.Contains(duration, "years")) {
		return CourseDuration{}, ErrDateDuration
	}

	return CourseDuration{
		value: value,
	}, nil
}

// String type converts the CourseDuration into string.
func (duration CourseDuration) String() string {
	return duration.value
}

// CourseRepository defines the  expected behaviour from a course storage.
type CourseRepository interface {
	Save(ctx context.Context, course Course) error
	FindAll() ([]Course, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

// Course is the data structure that represents a course.
type Course struct {
	id       CourseID
	name     CourseName
	duration CourseDuration
}

// NewCourse creates a new course.
func NewCourse(id, name, duration string) (Course, error) {
	idVO, err := NewCourseID(id)
	if err != nil {
		return Course{}, err
	}

	nameVO, err := NewCourseName(name)
	if err != nil {
		return Course{}, err
	}

	durationVO, err := NewCourseDuration(duration)
	if err != nil {
		return Course{}, err
	}

	return Course{
		id:       idVO,
		name:     nameVO,
		duration: durationVO,
	}, nil
}

// ID returns the course unique identifier.
func (c Course) ID() string {
	return c.id.value
}

// Name returns the course name.
func (c Course) Name() string {
	return c.name.value
}

// Duration returns the course duration.
func (c Course) Duration() string {
	return c.duration.value
}
