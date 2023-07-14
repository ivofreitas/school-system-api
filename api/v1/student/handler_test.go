package student

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/leantech/school-system-api/mock"
	"github.com/leantech/school-system-api/model"
	"github.com/stretchr/testify/assert"
)

var repoErr = errors.New("repo fail")

func TestCreate(t *testing.T) {
	testCases := []struct {
		Name          string
		Request       *model.Student
		CreateRepoErr error
		ExpectedError error
	}{
		{
			Name: "Test Case 1",
			Request: &model.Student{
				FirstName: "john",
				LastName:  "doe",
				SSID:      "123456",
			},
		},
		{
			Name: "Test Case 2",
			Request: &model.Student{
				FirstName: "john",
				LastName:  "doe",
				SSID:      "123456",
			},
			CreateRepoErr: repoErr,
			ExpectedError: repoErr,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()

			repositoryMock := &mock.StudentRepositoryMock{}

			repositoryMock.
				On("CreateStudent", ctx, tc.Request).
				Return(tc.CreateRepoErr)

			hdl := NewHandler(repositoryMock)
			response, err := hdl.Create(ctx, tc.Request)
			if tc.ExpectedError != nil {
				assert.Error(t, err)
				responseError := err.(*model.ResponseError)
				assert.Equal(t, tc.ExpectedError.Error(), responseError.DeveloperMessage)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, response)
		})
	}
}

func TestUpdate(t *testing.T) {
	testCases := []struct {
		Name          string
		Request       *model.UpdateStudentRequest
		GetRepoResult *model.Student
		GetRepoErr    error
		UpdateRepoErr error
		ExpectedError error
	}{
		{
			Name: "Test Case 1",
			Request: &model.UpdateStudentRequest{
				ID: "6090944a-6d75-40f6-91b3-8ff33406f768",
				Student: &model.Student{
					FirstName: "john",
					LastName:  "doe",
					SSID:      "123456",
				},
			},
			GetRepoResult: &model.Student{
				FirstName: "john",
				LastName:  "doe",
				SSID:      "123456",
			},
		},
		{
			Name: "Test Case 2",
			Request: &model.UpdateStudentRequest{
				ID: "6090944a-6d75-40f6-91b3-8ff33406f768",
				Student: &model.Student{
					FirstName: "john",
					LastName:  "doe",
					SSID:      "123456",
				},
			},
			GetRepoErr:    repoErr,
			ExpectedError: repoErr,
		},
		{
			Name: "Test Case 3",
			Request: &model.UpdateStudentRequest{
				ID: "6090944a-6d75-40f6-91b3-8ff33406f768",
				Student: &model.Student{
					FirstName: "john",
					LastName:  "doe",
					SSID:      "123456",
				},
			},
			GetRepoErr:    sql.ErrNoRows,
			ExpectedError: errors.New("Student not found"),
		},
		{
			Name: "Test Case 4",
			Request: &model.UpdateStudentRequest{
				ID: "6090944a-6d75-40f6-91b3-8ff33406f768",
				Student: &model.Student{
					FirstName: "john",
					LastName:  "doe",
					SSID:      "123456",
				},
			},
			UpdateRepoErr: repoErr,
			ExpectedError: repoErr,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()

			repositoryMock := &mock.StudentRepositoryMock{}

			repositoryMock.
				On("GetStudentByID", ctx, tc.Request.ID).
				Return(tc.GetRepoResult, tc.GetRepoErr)

			repositoryMock.
				On("UpdateStudent", ctx, tc.Request.ID, tc.Request.Student).
				Return(tc.UpdateRepoErr)

			hdl := NewHandler(repositoryMock)
			response, err := hdl.Update(ctx, tc.Request)
			if tc.ExpectedError != nil {
				assert.Error(t, err)
				responseError := err.(*model.ResponseError)
				assert.Equal(t, tc.ExpectedError.Error(), responseError.DeveloperMessage)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, response)
		})
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		Name          string
		Request       *model.DeleteStudentRequest
		GetRepoResult *model.Student
		GetRepoErr    error
		DeleteRepoErr error
		ExpectedError error
	}{
		{
			Name: "Test Case 1",
			Request: &model.DeleteStudentRequest{
				ID: "6090944a-6d75-40f6-91b3-8ff33406f768",
			},
			GetRepoResult: &model.Student{
				FirstName: "john",
				LastName:  "doe",
				SSID:      "123456",
			},
		},
		{
			Name: "Test Case 2",
			Request: &model.DeleteStudentRequest{
				ID: "6090944a-6d75-40f6-91b3-8ff33406f768",
			},
			GetRepoErr:    sql.ErrNoRows,
			ExpectedError: errors.New("Student not found"),
		},
		{
			Name: "Test Case 3",
			Request: &model.DeleteStudentRequest{
				ID: "6090944a-6d75-40f6-91b3-8ff33406f768",
			},
			GetRepoErr:    repoErr,
			ExpectedError: repoErr,
		},
		{
			Name: "Test Case 4",
			Request: &model.DeleteStudentRequest{
				ID: "6090944a-6d75-40f6-91b3-8ff33406f768",
			},
			DeleteRepoErr: repoErr,
			ExpectedError: repoErr,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()

			repositoryMock := &mock.StudentRepositoryMock{}

			repositoryMock.
				On("GetStudentByID", ctx, tc.Request.ID).
				Return(tc.GetRepoResult, tc.GetRepoErr)

			repositoryMock.
				On("DeleteStudent", ctx, tc.Request.ID).
				Return(tc.DeleteRepoErr)

			hdl := NewHandler(repositoryMock)
			response, err := hdl.Delete(ctx, tc.Request)
			if tc.ExpectedError != nil {
				assert.Error(t, err)
				responseError := err.(*model.ResponseError)
				assert.Equal(t, tc.ExpectedError.Error(), responseError.DeveloperMessage)
				return
			}

			assert.NoError(t, err)
			assert.Nil(t, response)
		})
	}
}
