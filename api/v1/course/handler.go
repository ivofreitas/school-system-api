package course

import (
	gocontext "context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/leantech/school-system-api/api/v1/student"
	"github.com/leantech/school-system-api/model"
	"golang.org/x/net/context"
)

type Repository interface {
	CreateCourse(ctx gocontext.Context, course *model.Course) error
	GetCourseByRoomID(ctx context.Context, roomID string) (*model.Course, error)
	EnrollStudents(ctx context.Context, students []string, courseID string) error
	student.Repository
}

type handler struct {
	repository Repository
}

func NewHandler(repository Repository) *handler {
	return &handler{repository: repository}
}

// Create
// @Summary create a course.
// @Param key body model.Course true "request body"
// @Tags course
// @Security Authorization
// @Accept json
// @Product json
// @Success 201 {object} model.Response{meta=model.Meta,records=[]model.Course}
// @Failure 400 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /course [post]
func (h *handler) Create(ctx gocontext.Context, param interface{}) (interface{}, error) {
	request := param.(*model.Course)
	request.ID = uuid.New().String()
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	if err := h.repository.CreateCourse(ctx, request); err != nil {
		return nil, model.ErrorDiscover(err)
	}

	return model.NewResponse(0, 0, 1, []interface{}{request}), nil
}

// EnrollStudent
// @Summary edit a course.
// @Param id path string true "course id"
// @Param course body model.EnrollStudentRequest true "course"
// @Tags course
// @Security Authorization
// @Accept json
// @Product json
// @Success 200 {object} model.Response{meta=model.Meta,records=[]model.EnrollStudentRequest}
// @Failure 404 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /course/enroll [post]
func (h *handler) EnrollStudent(ctx gocontext.Context, param interface{}) (interface{}, error) {
	request := param.(*model.EnrollStudentRequest)

	if students, err := h.repository.GetStudentsByIDs(ctx, request.Students); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrorDiscover(model.NotFound{DeveloperMessage: "Students not found"})
		}
		return nil, model.ErrorDiscover(err)
	} else {
		if len(students) != len(request.Students) {
			return nil, model.ErrorDiscover(model.NotFound{DeveloperMessage: "The requested students were not found"})
		}
	}

	course, err := h.repository.GetCourseByRoomID(ctx, request.RoomID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrorDiscover(model.NotFound{DeveloperMessage: "Course not found"})
		}

		return nil, model.ErrorDiscover(err)
	}

	if len(course.EnrolledStudents)+len(request.Students) > course.MaxStudentsNumber {
		return nil, model.ErrorDiscover(model.BadRequest{DeveloperMessage: "Course does not have enough capacity"})
	}

	if err := h.repository.EnrollStudents(ctx, request.Students, course.ID); err != nil {
		return nil, model.ErrorDiscover(err)
	}

	return model.NewResponse(0, 0, 1, []interface{}{request}), nil
}
