package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todoapp/model"
	"todoapp/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	mock.Mock
}

func (o mockService) CreateTodo(request service.CreateTodoRequest) (model.Todo, error) {
	args := o.Called(request)
	return args.Get(0).(model.Todo), args.Error(1)
}

func (o mockService) GetTodos() ([]model.Todo, error) {
	args := o.Called()
	return args.Get(0).([]model.Todo), args.Error(1)
}

func (o mockService) DeleteTodo(request service.DeleteTodoRequest) error {
	args := o.Called(request)
	return args.Error(0)
}

func Test_shouldCreateTodo(t *testing.T) {
	service := new(mockService)
	service.On("CreateTodo", mock.Anything).Return(model.Todo{}, nil)
	handler := NewTodoHandler(service)
	context, responseRecorder := getMockedDefaultContext()
	handler.CreateTodo(context)

	assert.True(t, len(context.Errors) == 0)
	assert.Equal(t, 201, responseRecorder.Code)
}

func Test_shouldGetTodoList(t *testing.T) {
	service := new(mockService)
	service.On("GetTodos", mock.Anything).Return([]model.Todo{model.Todo{Id: "id", Name: "todo 1", IsPending: false}}, nil)
	handler := NewTodoHandler(service)
	context, responseRecorder := getMockedDefaultContext()
	handler.GetTodoList(context)

	assert.True(t, len(context.Errors) == 0)
	assert.Equal(t, 200, responseRecorder.Code)
	assert.Equal(t, `[{"id":"id","name":"todo 1","isPending":false}]`, responseRecorder.Body.String())
}

func Test_shouldDeleteTodoList(t *testing.T) {
	service := new(mockService)
	service.On("DeleteTodo", mock.Anything).Return(nil)
	handler := NewTodoHandler(service)
	context, responseRecorder := getMockedDefaultContext()
	context.Params = []gin.Param{
		{Key: "todoId", Value: "abc123"},
	}
	handler.DeleteTodo(context)
	assert.True(t, len(context.Errors) == 0)
	assert.Equal(t, 200, responseRecorder.Code)
}

func getMockedDefaultContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	responseWriter := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(responseWriter)
	context.Request, _ = http.NewRequest("GET", "/", nil)

	query := map[string][]string{}
	context.Request.PostForm = query
	context.Request.URL.RawQuery = ""
	context.Params = []gin.Param{}

	return context, responseWriter

}
