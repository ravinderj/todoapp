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
	todo := model.NewTodo("todo 1")
	repository := NewTodoRepository(mongoProvider, "todo")
	err = repository.CreateTodo(todo)
	assert.Nil(t, err)
	var todoDao TodoDao
	collection := mongoProvider.Copy().DB("todo").C("todo")
	collection.Find(bson.M{}).One(&todoDao)
	assert.Equal(t, "todo 1", todoDao.Name)
	assert.Equal(t, false, todoDao.IsPending)
	collection.DropCollection()
	mongoProvider.Close()
}

func Test_shouldGetTodos(t *testing.T) {
	mongoProvider, err := mgo.Dial("localhost")
	todoDao := TodoDao{Id: bson.ObjectId("123456789012"), Name: "todo 2", IsPending: false}
	collection := mongoProvider.Copy().DB("todo").C("todo")
	err = collection.Insert(todoDao)
	repository := NewTodoRepository(mongoProvider, "todo")
	todos, err := repository.GetTodos()
	assert.Nil(t, err)
	assert.Equal(t, []TodoDao{todoDao}, todos)
	collection.DropCollection()
	mongoProvider.Close()
}
