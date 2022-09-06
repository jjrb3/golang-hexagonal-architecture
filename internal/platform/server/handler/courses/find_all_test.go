package courses

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	mooc "github.com/jjrb3/golang-hexagonal-architecture/internal"
	"github.com/jjrb3/golang-hexagonal-architecture/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_FindAll(t *testing.T) {
	t.Run("Get all available courses", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		r := gin.New()

		course, _ := mooc.NewCourse("8a1c5cdc-ba57-445a-994d-aa412d23723f", "Demo Course", "10 months")

		var courses = []mooc.Course{course}
		var expected = generateCoursesResponse(courses)

		courseRepository := new(storagemocks.CourseRepository)
		courseRepository.On("FindAll", mock.Anything).Return(courses, nil)

		r.GET("/courses", FindAllHandler(courseRepository))

		req, err := http.NewRequest(http.MethodGet, "/courses", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		var response []courseResponse
		if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
			log.Fatalln(err)
		}

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, expected, response)
	})

	t.Run("Get error by not found information", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		r := gin.New()

		var courses []mooc.Course

		courseRepository := new(storagemocks.CourseRepository)
		courseRepository.On("FindAll", mock.Anything).Return(courses, mooc.ErrEmptyDuration)

		r.GET("/courses", FindAllHandler(courseRepository))

		req, err := http.NewRequest(http.MethodGet, "/courses", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}

func generateCoursesResponse(courses []mooc.Course) []courseResponse {
	var res []courseResponse

	if len(courses) == 0 {
		return res
	}

	for _, course := range courses {
		res = append(res, courseResponse{
			ID:       course.ID().String(),
			Name:     course.Name().String(),
			Duration: course.Duration().String(),
		})
	}

	return res
}
