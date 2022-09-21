package courses

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	mooc "github.com/jjrb3/golang-hexagonal-architecture/internal"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/creating"
)

type createRequest struct {
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler(createCourseService creating.CourseService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":  err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}

		err := createCourseService.CreateCourse(ctx, uuid.New().String(), req.Name, req.Duration)

		if err != nil {
			switch {
			case errors.Is(err, mooc.ErrInvalidCourseID),
				errors.Is(err, mooc.ErrEmptyCourseName),
				errors.Is(err, mooc.ErrInvalidCourseID),
				errors.Is(err, mooc.ErrDateDuration):
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error":  err.Error(),
					"status": http.StatusBadRequest,
				})
				return
			default:
				fmt.Println(err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error":  err.Error(),
					"status": http.StatusInternalServerError,
				})
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
