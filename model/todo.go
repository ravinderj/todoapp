package model

type Todo struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	todoItems []TodoItem `json:"items"`
	isPending bool       `json:"isPending"`
}

type TodoItem struct {
	Id          int
	description string
}

func NewTodo(id int, name string) Todo {
	return Todo{
		Id:   id,
		Name: name,
	}
}

func NewTodoItem(id int, description string) TodoItem {
	return TodoItem{
		Id:          id,
		description: description,
	}
}

func (this *Todo) AddItem(item TodoItem) {
	this.todoItems = append(this.todoItems, item)
}

func (this *Todo) RemoveItem(id int) {
	index := findIndexById(this.todoItems, id)
	this.todoItems = append(this.todoItems[:index], this.todoItems[index+1:]...)
}

func findIndexById(todoItems []TodoItem, itemId int) int {
	for index, item := range todoItems {
		if item.Id == itemId {
			return index
		}
	}
	return 0
}
