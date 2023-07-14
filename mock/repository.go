package mock

import (
	"context"

	"github.com/leantech/school-system-api/model"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)

	var resultError error
	if args.Get(0) != nil {
		resultError = args.Get(0).(error)
	}

	return resultError
}

func (m *UserRepositoryMock) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	args := m.Called(ctx, username)

	var resultError error
	if args.Get(1) != nil {
		resultError = args.Get(1).(error)
	}

	return args.Get(0).(*model.User), resultError
}

type StudentRepositoryMock struct {
	mock.Mock
}

func (m *StudentRepositoryMock) CreateStudent(ctx context.Context, student *model.Student) error {
	args := m.Called(ctx, student)

	var resultError error
	if args.Get(0) != nil {
		resultError = args.Get(0).(error)
	}

	return resultError
}

func (m *StudentRepositoryMock) UpdateStudent(ctx context.Context, id string, student *model.Student) error {
	args := m.Called(ctx, id, student)

	var resultError error
	if args.Get(0) != nil {
		resultError = args.Get(0).(error)
	}

	return resultError
}

func (m *StudentRepositoryMock) GetStudentByID(ctx context.Context, id string) (*model.Student, error) {
	args := m.Called(ctx, id)

	var resultError error
	if args.Get(1) != nil {
		resultError = args.Get(1).(error)
	}

	return args.Get(0).(*model.Student), resultError
}

func (m *StudentRepositoryMock) GetStudentsByIDs(ctx context.Context, ids []string) ([]*model.Student, error) {
	args := m.Called(ctx, ids)

	var resultError error
	if args.Get(1) != nil {
		resultError = args.Get(1).(error)
	}

	return args.Get(0).([]*model.Student), resultError
}

func (m *StudentRepositoryMock) DeleteStudent(ctx context.Context, id string) error {
	args := m.Called(ctx, id)

	var resultError error
	if args.Get(0) != nil {
		resultError = args.Get(0).(error)
	}

	return resultError
}

type CourseRepositoryMock struct {
	mock.Mock
	StudentRepositoryMock
}

func (m *CourseRepositoryMock) CreateCourse(ctx context.Context, course *model.Course) error {
	args := m.Called(ctx, course)

	var resultError error
	if args.Get(0) != nil {
		resultError = args.Get(0).(error)
	}

	return resultError
}

func (m *CourseRepositoryMock) GetCourseByRoomID(ctx context.Context, roomID string) (*model.Course, error) {
	args := m.Called(ctx, roomID)

	var resultError error
	if args.Get(1) != nil {
		resultError = args.Get(1).(error)
	}

	return args.Get(0).(*model.Course), resultError
}

func (m *CourseRepositoryMock) EnrollStudents(ctx context.Context, students []string, courseID string) error {
	args := m.Called(ctx, students, courseID)

	var resultError error
	if args.Get(0) != nil {
		resultError = args.Get(0).(error)
	}

	return resultError
}
