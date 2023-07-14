package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/leantech/school-system-api/context"
	"github.com/leantech/school-system-api/log"
	"github.com/leantech/school-system-api/mock"
	"github.com/leantech/school-system-api/model"
	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"
)

var handlerErr = &model.ResponseError{DeveloperMessage: "handler fail"}

func TestHandleCreateStudent(t *testing.T) {
	testCases := []struct {
		Name           string
		Param          interface{}
		HandlerErr     error
		ExpectedStatus int
	}{
		{
			Name: "Test Case 1",
			Param: &model.Student{
				FirstName: "john",
				LastName:  "doe",
				SSID:      "123456",
			},
			ExpectedStatus: http.StatusCreated,
		},
		{
			Name:           "Test Case 2",
			Param:          "",
			ExpectedStatus: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			b, _ := json.Marshal(&tc.Param)
			req := httptest.NewRequest(http.MethodPost, "/v1/student", strings.NewReader(string(b)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			e := echo.New()
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			ctx := log.InitParams(c.Request().Context())
			ctx = context.Set(ctx, "username", "joe.doe")
			ctx = context.Set(ctx, "role", "admin")
			c.SetRequest(c.Request().WithContext(ctx))

			e.Binder = NewBinder()
			e.Validator = NewValidator()

			handlerMock := &mock.StudentHandlerMock{}
			handlerMock.
				On("Create", ctx, testifymock.Anything).
				Return(tc.Param, tc.HandlerErr)

			ctrl := NewController(handlerMock.Create, http.StatusCreated, new(model.Student))
			ctrl.Handle(c)

			assert.Equal(t, tc.ExpectedStatus, rec.Code)
			assert.NotNil(t, rec.Body)
		})
	}
}

func TestHandleUpdateStudent(t *testing.T) {
	testCases := []struct {
		Name           string
		Param          interface{}
		HandlerErr     error
		ExpectedError  error
		ExpectedStatus int
	}{
		{
			Name: "Test Case 1",
			Param: &model.UpdateStudentRequest{
				ID: "6090944a-6d75-40f6-91b3-8ff33406f768",
				Student: &model.Student{
					FirstName: "john",
					LastName:  "doe",
					SSID:      "123456",
				},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "Test Case 2",
			Param:          "",
			ExpectedStatus: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			b, _ := json.Marshal(&tc.Param)
			req := httptest.NewRequest(http.MethodPut, "/v1/student", strings.NewReader(string(b)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			e := echo.New()
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			ctx := log.InitParams(c.Request().Context())
			ctx = context.Set(ctx, "username", "joe.doe")
			ctx = context.Set(ctx, "role", "admin")
			c.SetRequest(c.Request().WithContext(ctx))

			e.Binder = NewBinder()
			e.Validator = NewValidator()

			handlerMock := &mock.StudentHandlerMock{}
			handlerMock.
				On("Update", ctx, testifymock.Anything).
				Return(tc.Param, tc.HandlerErr)

			ctrl := NewController(handlerMock.Update, http.StatusOK, new(model.UpdateStudentRequest))
			err := ctrl.Handle(c)
			if tc.ExpectedError != nil {
				assert.Error(t, err)
				responseError := err.(*model.ResponseError)
				assert.Equal(t, tc.ExpectedError.Error(), responseError.DeveloperMessage)
				return
			}

			assert.Equal(t, tc.ExpectedStatus, rec.Code)
			assert.NotNil(t, rec.Body)
			assert.NoError(t, err)
		})
	}
}

func TestHandleDeleteStudent(t *testing.T) {
	testCases := []struct {
		Name           string
		Param          interface{}
		HandlerErr     error
		ExpectedError  error
		ExpectedStatus int
	}{
		{
			Name: "Test Case 1",
			Param: &model.DeleteStudentRequest{
				ID: "6090944a-6d75-40f6-91b3-8ff33406f768",
			},
			ExpectedStatus: http.StatusNoContent,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodDelete, "/v1/student", nil)

			e := echo.New()
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			ctx := log.InitParams(c.Request().Context())
			ctx = context.Set(ctx, "username", "joe.doe")
			ctx = context.Set(ctx, "role", "admin")
			c.SetRequest(c.Request().WithContext(ctx))

			e.Binder = NewBinder()
			e.Validator = NewValidator()

			handlerMock := &mock.StudentHandlerMock{}
			handlerMock.
				On("Delete", ctx, testifymock.Anything).
				Return(tc.Param, tc.HandlerErr)

			ctrl := NewController(handlerMock.Delete, http.StatusNoContent, new(model.DeleteStudentRequest))
			err := ctrl.Handle(c)
			if tc.ExpectedError != nil {
				assert.Error(t, err)
				responseError := err.(*model.ResponseError)
				assert.Equal(t, tc.ExpectedError.Error(), responseError.DeveloperMessage)
				return
			}

			assert.Equal(t, tc.ExpectedStatus, rec.Code)
			assert.NotNil(t, rec.Body)
			assert.NoError(t, err)
		})
	}
}
