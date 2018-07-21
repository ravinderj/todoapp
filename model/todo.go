package model

type Todo struct {
	// Id   int    `json:"id"`
	Name string `json:"name"`
	// TodoItems []TodoItem `bson:"items"`
	IsPending bool `json:"isPending"`
}

// type TodoItem struct {
// 	Id          int
// 	description string
// }

func NewTodo(name string) Todo {
	return Todo{
		Name:      name,
		IsPending: false,
	}
}

// func NewTodoItem(id int, description string) TodoItem {
// 	return TodoItem{
// 		Id:          id,
// 		description: description,
// 	}
// }

// func (this *Todo) AddItem(item TodoItem) {
// 	this.TodoItems = append(this.TodoItems, item)
// }

// func (this *Todo) RemoveItem(id int) {
// 	index := findIndexById(this.TodoItems, id)
// 	this.TodoItems = append(this.TodoItems[:index], this.TodoItems[index+1:]...)
// }

// func findIndexById(todoItems []TodoItem, itemId int) int {
// 	for index, item := range todoItems {
// 		if item.Id == itemId {
// 			return index
// 		}
// 	}
// 	return 0
// }
