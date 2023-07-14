package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type StudentHandlerMock struct {
	mock.Mock
}

func (m *StudentHandlerMock) Create(ctx context.Context, param interface{}) (interface{}, error) {
	args := m.Called(ctx, param)

	var resultError error
	if args.Get(1) != nil {
		resultError = args.Get(1).(error)
	}

	return args.Get(0), resultError
}

func (m *StudentHandlerMock) Update(ctx context.Context, param interface{}) (interface{}, error) {
	args := m.Called(ctx, param)

	var resultError error
	if args.Get(1) != nil {
		resultError = args.Get(1).(error)
	}

	return args.Get(0), resultError
}

func (m *StudentHandlerMock) Delete(ctx context.Context, param interface{}) (interface{}, error) {
	args := m.Called(ctx, param)

	var resultError error
	if args.Get(1) != nil {
		resultError = args.Get(1).(error)
	}

	return args.Get(0), resultError
}

type UserHandlerMock struct {
	mock.Mock
}

func (m *UserHandlerMock) Create(ctx context.Context, param interface{}) (interface{}, error) {
	args := m.Called(ctx, param)

	var resultError error
	if args.Get(1) != nil {
		resultError = args.Get(1).(error)
	}

	return args.Get(0), resultError
}

func (m *UserHandlerMock) Login(ctx context.Context, param interface{}) (interface{}, error) {
	args := m.Called(ctx, param)

	var resultError error
	if args.Get(1) != nil {
		resultError = args.Get(1).(error)
	}

	return args.Get(0), resultError
}

type CourseHandlerMock struct {
	mock.Mock
}

func (m *CourseHandlerMock) Create(ctx context.Context, param interface{}) (interface{}, error) {
	args := m.Called(ctx, param)

	var resultError error
	if args.Get(1) != nil {
		resultError = args.Get(1).(error)
	}

	return args.Get(0), resultError
}

func (m *CourseHandlerMock) EnrollStudent(ctx context.Context, param interface{}) (interface{}, error) {
	args := m.Called(ctx, param)

	var resultError error
	if args.Get(1) != nil {
		resultError = args.Get(1).(error)
	}

	return args.Get(0), resultError
}
