package handler

import (
	"errors"
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
	createTodoResponse, err := this.todoService.CreateTodo(request)
	if err != nil {
		context.Error(err)
	}
	context.JSON(http.StatusCreated, createTodoResponse)
}

func (this *todoHandler) DeleteTodo(context *gin.Context) {
	request := this.getDeleteTodoRequest(context)
	if len(context.Errors) > 0 {
		return
	}
	err := this.todoService.DeleteTodo(request)
	if err != nil {
		context.Error(err)
	}
	context.Status(http.StatusOK)
}

func (this *todoHandler) getDeleteTodoRequest(context *gin.Context) service.DeleteTodoRequest {
	request := service.DeleteTodoRequest{}
	todoId, hasValue := context.Params.Get("todoId")
	if !hasValue || todoId == "" {
		context.Error(errors.New("Missing todo reference Id"))
		return request
	}
	request.TodoId = todoId
	return request
}
