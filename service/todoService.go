package service

import (
	"todoapp/model"
	"todoapp/repository"

	"gopkg.in/mgo.v2/bson"
)

type TodoService interface {
	CreateTodo(request CreateTodoRequest) error
	DeleteTodo(request DeleteTodoRequest) error
	GetTodos() ([]model.Todo, error)
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) *todoService {
	return &todoService{
		todoRepository: todoRepository,
	}
}

func (this *todoService) CreateTodo(request CreateTodoRequest) error {
	todo, err := this.createTodoFromRequest(request)
	if err != nil {
		return err
	}
	err = this.todoRepository.CreateTodo(todo)
	if err != nil {
		return err
	}
	return nil
}

func (this *todoService) DeleteTodo(request DeleteTodoRequest) error {
	todoId := request.TodoId
	err := this.todoRepository.DropTodo(todoId)
	return err
}

func (this *todoService) GetTodos() ([]model.Todo, error) {
	todosDao, err := this.todoRepository.GetTodos()
	if err != nil {
		return nil, err
	}
	todos := this.mapTodoDaoToModel(todosDao)
	return todos, nil
}

func (this *todoService) mapTodoDaoToModel(todosDao []repository.TodoDao) []model.Todo {
	var todos []model.Todo
	for _, todoDao := range todosDao {
		todos = append(todos, model.Todo{Id: todoDao.Id.Hex(), Name: todoDao.Name, IsPending: todoDao.IsPending})
	}
	return todos
}

func (this *todoService) createTodoFromRequest(request CreateTodoRequest) (model.Todo, error) {
	todoId := bson.NewObjectId().Hex()
	return model.NewTodo(request.Name, todoId), nil
}
