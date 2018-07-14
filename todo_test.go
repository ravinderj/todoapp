package todo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shouldAddItemToTodo(t *testing.T) {
	todo := NewTodo(1, "Home Work")
	todoItem := NewTodoItem(1, "science note book")
	todo.AddItem(todoItem)
	assert.Equal(t, 1, len(todo.todoItems))
}

func Test_shouldRemoveItemFromTodo(t *testing.T) {
	todo := NewTodo(1, "Home Work")
	todoItem := NewTodoItem(1, "science note book")
	todo.AddItem(todoItem)
	todo.RemoveItem(1)
	assert.Equal(t, 0, len(todo.todoItems))
}
