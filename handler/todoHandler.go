package handler

import (
	"net/http"
	"todoapp/service"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(service service.TodoService) todoHandler {
	return todoHandler{
		todoService: service,
	}
}

func (this *todoHandler) GetTodoList(context *gin.Context) {
	context.JSON(http.StatusOK, nil)
}

func (this *todoHandler) CreateTodo(context *gin.Context) {
	request := createTodoRequest{}
	context.Bind(&request)
	response, err := this.todoService.CreateTodo(request)
	if err != nil {
		context.Error(err)
	}
	context.JSON(http.StatusCreated, response)
}
