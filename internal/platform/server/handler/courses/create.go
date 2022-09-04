package courses

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	mooc "github.com/jjrb3/golang-hexagonal-architecture/internal"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler(courseRepository mooc.CourseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course, err := mooc.NewCourse(req.ID, req.Name, req.Duration)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err = courseRepository.Save(ctx, course); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
