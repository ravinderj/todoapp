package service

import (
	"testing"
	"todoapp/model"
	"todoapp/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/mgo.v2/bson"
)

type mockRepository struct {
	mock.Mock
}

func (o mockRepository) CreateTodo(todo model.Todo) error {
	args := o.Called(todo)
	return args.Error(0)
}

func (o mockRepository) GetTodos() ([]repository.TodoDao, error) {
	args := o.Called()
	return args.Get(0).([]repository.TodoDao), args.Error(1)
}

func (o mockRepository) DropTodo(todoId string) error {
	args := o.Called()
	return args.Error(0)
}

func Test_shouldCreateTodo(t *testing.T) {
	repository := new(mockRepository)
	repository.On("CreateTodo", mock.Anything).Return(nil)
	service := NewTodoService(repository)
	createTodoRequest := CreateTodoRequest{Name: "todo 1"}
	todo, err := service.CreateTodo(createTodoRequest)
	assert.Nil(t, err)
	assert.Equal(t, todo.Name, "todo 1")
}

func Test_shouldDeleteTodo(t *testing.T) {
	repository := new(mockRepository)
	repository.On("DropTodo", mock.Anything).Return(nil)
	service := NewTodoService(repository)
	deleteTodoRequest := DeleteTodoRequest{TodoId: "todoId"}
	err := service.DeleteTodo(deleteTodoRequest)
	assert.Nil(t, err)
}

func Test_shouldGetTodos(t *testing.T) {
	mockRepository := new(mockRepository)
	todoDao := repository.TodoDao{
		Id:        bson.ObjectId("12345"),
		Name:      "todo 1",
		IsPending: true,
	}
	todo := model.NewTodo("todo 1", "3132333435")
	mockRepository.On("GetTodos", mock.Anything).Return([]repository.TodoDao{todoDao}, nil)
	service := NewTodoService(mockRepository)
	todos, err := service.GetTodos()
	assert.Nil(t, err)
	assert.Equal(t, []model.Todo{todo}, todos)
}
