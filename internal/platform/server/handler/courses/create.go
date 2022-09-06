package courses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	mooc "github.com/jjrb3/golang-hexagonal-architecture/internal"
)

type createRequest struct {
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler(courseRepository mooc.CourseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":  err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}

		course, err := mooc.NewCourse(uuid.New().String(), req.Name, req.Duration)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":  err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}

		if err = courseRepository.Save(ctx, course); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":  err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
