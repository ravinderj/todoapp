package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todoapp/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_shouldCreateTodo(t *testing.T) {
	todos := []model.Todo{model.NewTodo(1, "todo 1")}
	handler := NewTodoHandler(todos)
	context, responseRecorder := getMockedDefaultContext()
	context.Request = httptest.NewRequest("POST", "http://localhost/todos", strings.NewReader(`{"name":"New Todo"}`))
	context.Request.Header.Set("Content-Type", "application/json")
	handler.CreateTodo(context)

	assert.True(t, len(context.Errors) == 0)

	assert.Equal(t, 202, responseRecorder.Code)
	assert.Equal(t, `{"id":1,"name":"New Todo"}`, responseRecorder.Body.String())
}

func Test_shouldGetTodoList(t *testing.T) {
	todos := []model.Todo{model.NewTodo(1, "todo 1")}
	handler := NewTodoHandler(todos)
	context, responseRecorder := getMockedDefaultContext()
	context.Request = httptest.NewRequest("GET", "http://localhost/todos/list", nil)
	handler.GetTodoList(context)

	assert.True(t, len(context.Errors) == 0)

	assert.Equal(t, 200, responseRecorder.Code)
	assert.Equal(t, `[{"id":1,"name":"todo 1"}]`, responseRecorder.Body.String())
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
