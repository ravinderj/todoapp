package model

type Todo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	// TodoItems []TodoItem `bson:"items"`
	IsPending bool `json:"isPending"`
}

func NewTodo(name string, id string) Todo {
	return Todo{
		Id:        id,
		Name:      name,
		IsPending: false,
	}
}
