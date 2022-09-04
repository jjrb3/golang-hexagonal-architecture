// Code generated by mockery v2.12.2. DO NOT EDIT.

package storagemocks

import (
	context "context"

	mooc "github.com/jjrb3/golang-hexagonal-architecture/internal"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// CourseRepository is an autogenerated mock type for the CourseRepository type
type CourseRepository struct {
	mock.Mock
}

// FindAll provides a mock function with given fields:
func (_m *CourseRepository) FindAll() ([]mooc.Course, error) {
	ret := _m.Called()

	var r0 []mooc.Course
	if rf, ok := ret.Get(0).(func() []mooc.Course); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mooc.Course)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, course
func (_m *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	ret := _m.Called(ctx, course)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, mooc.Course) error); ok {
		r0 = rf(ctx, course)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCourseRepository creates a new instance of CourseRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCourseRepository(t testing.TB) *CourseRepository {
	mock := &CourseRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
