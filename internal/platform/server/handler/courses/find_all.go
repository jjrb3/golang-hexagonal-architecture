package courses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mooc "github.com/jjrb3/golang-hexagonal-architecture/internal"
)

type courseResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Duration string `json:"duration"`
}

// FindAllHandler returns an HTTP handler with all courses.
func FindAllHandler(courseRepository mooc.CourseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		courses, err := courseRepository.FindAll()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		cr := make([]courseResponse, 0, len(courses))

		for _, course := range courses {
			cr = append(cr, courseResponse{
				ID:       course.ID(),
				Name:     course.Name(),
				Duration: course.Duration(),
			})
		}

		ctx.JSON(http.StatusOK, cr)
	}
}
