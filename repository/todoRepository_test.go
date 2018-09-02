package repository

import (
	"testing"
	"todoapp/model"

	"github.com/stretchr/testify/assert"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Test_shouldCreateTodo(t *testing.T) {
	mongoProvider, err := mgo.Dial("localhost")
	collection := mongoProvider.Copy().DB("todo").C("todo")
	collection.DropCollection()

	todo := model.NewTodo("todo 1", bson.ObjectId("123456789012").Hex())
	repository := NewTodoRepository(mongoProvider, "todo")
	err = repository.CreateTodo(todo)
	assert.Nil(t, err)

	var todoDao TodoDao
	collection.Find(bson.M{}).One(&todoDao)
	assert.Equal(t, todoDao, TodoDao{Id: bson.ObjectId("123456789012"), Name: "todo 1", IsPending: true})
	mongoProvider.Close()
}

func Test_shouldDeleteTodo(t *testing.T) {
	mongoProvider, _ := mgo.Dial("localhost")
	collection := mongoProvider.Copy().DB("todo").C("todo")
	collection.DropCollection()

	objectID := bson.ObjectId("123456789012")
	todoDao := TodoDao{Id: objectID, Name: "todo 2", IsPending: false}
	collection.Insert(todoDao)

	repository := NewTodoRepository(mongoProvider, "todo")
	repository.DropTodo(objectID.Hex())
	var todos []TodoDao
	collection.Find(bson.M{}).All(&todos)
	assert.Equal(t, 0, len(todos))
	mongoProvider.Close()
}

func Test_shouldGetTodos(t *testing.T) {
	mongoProvider, err := mgo.Dial("localhost")
	collection := mongoProvider.Copy().DB("todo").C("todo")
	collection.DropCollection()

	todoDao := TodoDao{Id: bson.ObjectId("123456789012"), Name: "todo 2", IsPending: false}
	err = collection.Insert(todoDao)

	repository := NewTodoRepository(mongoProvider, "todo")
	todos, err := repository.GetTodos()
	assert.Nil(t, err)
	assert.Equal(t, []TodoDao{todoDao}, todos)
	mongoProvider.Close()
}
