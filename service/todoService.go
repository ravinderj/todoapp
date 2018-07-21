package service

type TodoService interface {
	CreateTodo(request CreateTodoRequest) (CreateTodoResponse, err)
}

type todoService struct {
	todoRepository model.TodoRepository
}

func (this *todoService) CreateTodo(request CreateTodoRequest) err {
	todo, err := this.createTodoFromRequest(request)
	if err != nil {
		return err
	}
	// err := this.todoRepository.CreateTodo(todo)
	// if err != nil {
	// 	return err
	// }
	return nil
}
