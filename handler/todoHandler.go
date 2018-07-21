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
	todos, err := this.todoService.GetTodos()
	if err != nil {
		context.Error(err)
	}
	context.JSON(http.StatusOK, todos)
}

func (this *todoHandler) CreateTodo(context *gin.Context) {
	request := service.CreateTodoRequest{}
	context.Bind(&request)
	err := this.todoService.CreateTodo(request)
	if err != nil {
		context.Error(err)
	}
	context.JSON(http.StatusCreated, nil)
}
