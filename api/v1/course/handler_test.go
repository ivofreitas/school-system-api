package course

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
		Request       *model.Course
		CreateRepoErr error
		ExpectedError error
	}{
		{
			Name: "Test Case 1",
			Request: &model.Course{
				RoomID:            "0001",
				MaxStudentsNumber: 10,
				CourseName:        "statistics",
			},
		},
		{
			Name: "Test Case 2",
			Request: &model.Course{
				RoomID:            "0001",
				MaxStudentsNumber: 10,
				CourseName:        "statistics",
			},
			CreateRepoErr: repoErr,
			ExpectedError: repoErr,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()

			repositoryMock := &mock.CourseRepositoryMock{}

			repositoryMock.
				On("CreateCourse", ctx, tc.Request).
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

func TestEnrollStudent(t *testing.T) {
	testCases := []struct {
		Name              string
		Request           *model.EnrollStudentRequest
		GetStudentResult  []*model.Student
		GetStudentErr     error
		GetCourseResult   *model.Course
		GetCourseErr      error
		EnrollStudentsErr error
		ExpectedError     error
	}{
		{
			Name: "Test Case 1",
			Request: &model.EnrollStudentRequest{
				RoomID:   "0001",
				Students: []string{"5d03c45c-212b-4033-915a-fdfede91a0e9"},
			},
			GetStudentResult: []*model.Student{{
				FirstName: "joe",
				LastName:  "doe",
				SSID:      "123456",
			}},
			GetCourseResult: &model.Course{
				ID:                "92e3babc-abb3-4899-b07b-8e070edbf53c",
				RoomID:            "0001",
				MaxStudentsNumber: 10,
				CourseName:        "statistics",
			},
		},
		{
			Name: "Test Case 2",
			Request: &model.EnrollStudentRequest{
				RoomID:   "0001",
				Students: []string{"5d03c45c-212b-4033-915a-fdfede91a0e9"},
			},
			GetStudentResult: []*model.Student{{
				FirstName: "joe",
				LastName:  "doe",
				SSID:      "123456",
			}},
			GetCourseResult: &model.Course{
				ID:                "92e3babc-abb3-4899-b07b-8e070edbf53c",
				RoomID:            "0001",
				MaxStudentsNumber: 10,
				CourseName:        "statistics",
			},
			GetStudentErr: repoErr,
			ExpectedError: repoErr,
		},
		{
			Name: "Test Case 3",
			Request: &model.EnrollStudentRequest{
				RoomID:   "0001",
				Students: []string{"5d03c45c-212b-4033-915a-fdfede91a0e9"},
			},
			GetStudentResult: []*model.Student{{
				FirstName: "joe",
				LastName:  "doe",
				SSID:      "123456",
			}},
			GetCourseResult: &model.Course{
				ID:                "92e3babc-abb3-4899-b07b-8e070edbf53c",
				RoomID:            "0001",
				MaxStudentsNumber: 10,
				CourseName:        "statistics",
			},
			GetStudentErr: sql.ErrNoRows,
			ExpectedError: errors.New("Students not found"),
		},
		{
			Name: "Test Case 4",
			Request: &model.EnrollStudentRequest{
				RoomID:   "0001",
				Students: []string{"5d03c45c-212b-4033-915a-fdfede91a0e9"},
			},
			GetStudentResult: []*model.Student{{
				FirstName: "joe",
				LastName:  "doe",
				SSID:      "123456",
			}, {
				FirstName: "mary",
				LastName:  "jane",
				SSID:      "124",
			}},
			GetCourseResult: &model.Course{
				ID:                "92e3babc-abb3-4899-b07b-8e070edbf53c",
				RoomID:            "0001",
				MaxStudentsNumber: 10,
				CourseName:        "statistics",
			},
			ExpectedError: errors.New("The requested students were not found"),
		},
		{
			Name: "Test Case 5",
			Request: &model.EnrollStudentRequest{
				RoomID:   "0001",
				Students: []string{"5d03c45c-212b-4033-915a-fdfede91a0e9"},
			},
			GetStudentResult: []*model.Student{{
				FirstName: "joe",
				LastName:  "doe",
				SSID:      "123456",
			}},
			GetCourseResult: &model.Course{
				ID:                "92e3babc-abb3-4899-b07b-8e070edbf53c",
				RoomID:            "0001",
				MaxStudentsNumber: 10,
				CourseName:        "statistics",
			},
			GetCourseErr:  sql.ErrNoRows,
			ExpectedError: errors.New("Course not found"),
		},
		{
			Name: "Test Case 6",
			Request: &model.EnrollStudentRequest{
				RoomID:   "0001",
				Students: []string{"5d03c45c-212b-4033-915a-fdfede91a0e9"},
			},
			GetStudentResult: []*model.Student{{
				FirstName: "joe",
				LastName:  "doe",
				SSID:      "123456",
			}},
			GetCourseResult: &model.Course{
				ID:                "92e3babc-abb3-4899-b07b-8e070edbf53c",
				RoomID:            "0001",
				MaxStudentsNumber: 10,
				CourseName:        "statistics",
			},
			GetCourseErr:  repoErr,
			ExpectedError: repoErr,
		},
		{
			Name: "Test Case 7",
			Request: &model.EnrollStudentRequest{
				RoomID:   "0001",
				Students: []string{"5d03c45c-212b-4033-915a-fdfede91a0e9"},
			},
			GetStudentResult: []*model.Student{{
				FirstName: "joe",
				LastName:  "doe",
				SSID:      "123456",
			}},
			GetCourseResult: &model.Course{
				ID:                "92e3babc-abb3-4899-b07b-8e070edbf53c",
				RoomID:            "0001",
				MaxStudentsNumber: 0,
				CourseName:        "statistics",
			},
			ExpectedError: errors.New("Course does not have enough capacity"),
		},
		{
			Name: "Test Case 8",
			Request: &model.EnrollStudentRequest{
				RoomID:   "0001",
				Students: []string{"5d03c45c-212b-4033-915a-fdfede91a0e9"},
			},
			GetStudentResult: []*model.Student{{
				FirstName: "joe",
				LastName:  "doe",
				SSID:      "123456",
			}},
			GetCourseResult: &model.Course{
				ID:                "92e3babc-abb3-4899-b07b-8e070edbf53c",
				RoomID:            "0001",
				MaxStudentsNumber: 10,
				CourseName:        "statistics",
			},
			EnrollStudentsErr: repoErr,
			ExpectedError:     repoErr,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()

			repositoryMock := &mock.CourseRepositoryMock{}

			repositoryMock.StudentRepositoryMock.
				On("GetStudentsByIDs", ctx, tc.Request.Students).
				Return(tc.GetStudentResult, tc.GetStudentErr)

			repositoryMock.
				On("GetCourseByRoomID", ctx, tc.Request.RoomID).
				Return(tc.GetCourseResult, tc.GetCourseErr)

			repositoryMock.
				On("EnrollStudents", ctx, tc.Request.Students, tc.GetCourseResult.ID).
				Return(tc.EnrollStudentsErr)

			hdl := NewHandler(repositoryMock)
			response, err := hdl.EnrollStudent(ctx, tc.Request)
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
