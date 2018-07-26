package app

import (
	"todoapp/handler"
	"todoapp/repository"
	"todoapp/service"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func SetupRouter() *gin.Engine {
	mongoProvider, _ := mgo.Dial("localhost")
	todoRepository := repository.NewTodoRepository(mongoProvider, "todo")
	todoService := service.NewTodoService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoService)
	router := gin.Default()
	router.GET("/", serveHome)
	router.GET("todos/list", todoHandler.GetTodoList)
	router.POST("todo", todoHandler.CreateTodo)
	router.DELETE("todo/:todoId", todoHandler.DeleteTodo)
	return router
}

func serveHome(context *gin.Context) {
	context.String(200, "Welcome to Home")
}

func Start() {
	router := SetupRouter()
	router.Run(":8080")
}
