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
