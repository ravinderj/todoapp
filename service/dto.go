package service

type CreateTodoRequest struct {
	Name string `form:"name"`
}

type DeleteTodoRequest struct {
	TodoId string
}
