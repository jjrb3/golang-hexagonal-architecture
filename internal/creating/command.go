package creating

import (
	"context"
	"errors"

	"github.com/jjrb3/golang-hexagonal-architecture/kit/command"
)

const CourseCommandType command.Type = "command.creating.course"

// CourseCommand is the command dispatched to create a new course.
type CourseCommand struct {
	id       string
	name     string
	duration string
}

// NewCourseCommand creates a new CourseCommand.
func NewCourseCommand(id, name, duration string) CourseCommand {
	return CourseCommand{
		id:       id,
		name:     name,
		duration: duration,
	}
}

func (c CourseCommand) Type() command.Type {
	return CourseCommandType
}

// CourseCommandHandler is the command handler responsible for creating courses.
type CourseCommandHandler struct {
	service CourseService
}

// NewCourseCommanderHandler initializes a new CourseCommandHandler.
func NewCourseCommanderHandler(service CourseService) CourseCommandHandler {
	return CourseCommandHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (h CourseCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	createCourseCMD, ok := cmd.(CourseCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.CreateCourse(
		ctx,
		createCourseCMD.id,
		createCourseCMD.name,
		createCourseCMD.duration,
	)
}
