package model

type Todo struct {
	// Id   int    `json:"id"`
	Name string `json:"name"`
	// TodoItems []TodoItem `bson:"items"`
	IsPending bool `json:"isPending"`
}

func NewTodo(name string) Todo {
	return Todo{
		Name:      name,
		IsPending: false,
	}
}
