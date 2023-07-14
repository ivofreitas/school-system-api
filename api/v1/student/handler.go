package student

import (
	gocontext "context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/leantech/school-system-api/model"
)

type Repository interface {
	CreateStudent(ctx gocontext.Context, student *model.Student) error
	UpdateStudent(ctx gocontext.Context, id string, student *model.Student) error
	GetStudentByID(ctx gocontext.Context, id string) (*model.Student, error)
	GetStudentsByIDs(ctx gocontext.Context, ids []string) ([]*model.Student, error)
	DeleteStudent(ctx gocontext.Context, id string) error
}

type handler struct {
	repository Repository
}

func NewHandler(repository Repository) *handler {
	return &handler{repository: repository}
}

// Create
// @Summary create a student.
// @Param key body model.Student true "request body"
// @Tags student
// @Security Authorization
// @Accept json
// @Product json
// @Success 201 {object} model.Response{meta=model.Meta,records=[]model.Student}
// @Failure 400 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /student [post]
func (h *handler) Create(ctx gocontext.Context, param interface{}) (interface{}, error) {
	request := param.(*model.Student)
	request.ID = uuid.New().String()
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	if err := h.repository.CreateStudent(ctx, request); err != nil {
		return nil, model.ErrorDiscover(err)
	}

	return model.NewResponse(0, 0, 1, []interface{}{request}), nil
}

// Update
// @Summary edit a student.
// @Param id path string true "student id"
// @Param student body model.Student true "student"
// @Tags student
// @Security Authorization
// @Accept json
// @Product json
// @Success 200 {object} model.Response{meta=model.Meta,records=[]model.Student}
// @Failure 400 {object} model.ResponseError
// @Failure 404 {object} model.ResponseError "No student found"
// @Failure 500 {object} model.ResponseError
// @Router /student/{id} [put]
func (h *handler) Update(ctx gocontext.Context, param interface{}) (interface{}, error) {
	request := param.(*model.UpdateStudentRequest)

	if _, err := h.repository.GetStudentByID(ctx, request.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrorDiscover(model.NotFound{DeveloperMessage: "Student not found"})
		}
		return nil, model.ErrorDiscover(err)
	}

	request.Student.UpdatedAt = time.Now()
	if err := h.repository.UpdateStudent(ctx, request.ID, request.Student); err != nil {
		return nil, model.ErrorDiscover(err)
	}

	return model.NewResponse(0, 0, 1, []interface{}{request}), nil
}

// Delete
// @Summary delete a student.
// @Param id path string true "student id"
// @Tags student
// @Security Authorization
// @Accept json
// @Product json
// @Success 200 {object} model.Response{meta=model.Meta,records=[]model.Student}
// @Failure 400 {object} model.ResponseError
// @Failure 404 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /student/{id} [delete]
func (h *handler) Delete(ctx gocontext.Context, param interface{}) (interface{}, error) {
	request := param.(*model.DeleteStudentRequest)

	if _, err := h.repository.GetStudentByID(ctx, request.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrorDiscover(model.NotFound{DeveloperMessage: "Student not found"})
		}
		return nil, model.ErrorDiscover(err)
	}

	if err := h.repository.DeleteStudent(ctx, request.ID); err != nil {
		return nil, model.ErrorDiscover(err)
	}

	return nil, nil
}
