package app

import (
	"todoapp/handler"
	"todoapp/model"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	var todos []model.Todo
	todoHandler := handler.NewTodoHandler(todos)
	router := gin.Default()
	router.GET("/", serveHome)
	router.GET("todos/list", todoHandler.GetTodoList)
	router.POST("todo", todoHandler.CreateTodo)
	return router
}

func serveHome(context *gin.Context) {
	context.String(200, "Welcome to Home")
}

func Start() {
	router := SetupRouter()
	router.Run(":8080")
}
