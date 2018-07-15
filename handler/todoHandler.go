package handler

import (
	"todoapp/model"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	Todos []model.Todo
}

func NewTodoHandler(todos []model.Todo) todoHandler {
	return todoHandler{
		Todos: todos,
	}
}

func (this *todoHandler) GetTodoList(context *gin.Context) {
	context.JSON(200, this.Todos)
}

func (this *todoHandler) CreateTodo(context *gin.Context) {
	request := createTodoRequest{}
	context.Bind(&request)
	todo := model.NewTodo(1, request.Name)
	this.Todos = append(this.Todos, todo)
	context.JSON(202, todo)
}

type createTodoRequest struct {
	Name string `form:"name"`
}
