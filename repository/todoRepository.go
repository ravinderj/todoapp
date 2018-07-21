package repository

import (
	"todoapp/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TodoRepository interface {
	CreateTodo(todo model.Todo) error
	GetTodos() ([]TodoDao, error)
}

type todoRepository struct {
	provider *mgo.Session
}

func (this *todoRepository) CreateTodo(todo model.Todo) error {
	session := this.provider.Copy()
	defer session.Close()
	todoCollection := session.DB("todo").C("todo")
	objectId := bson.NewObjectId()
	err := todoCollection.Insert(TodoDao{
		Id:        objectId,
		Name:      todo.Name,
		IsPending: todo.IsPending,
	})
	if err != nil {
		return err
	}
	return nil
}

func (this *todoRepository) GetTodos() ([]TodoDao, error) {
	session := this.provider.Copy()
	defer session.Close()
	todoCollection := session.DB("todo").C("todo")
	var todos []TodoDao
	err := todoCollection.Find(bson.M{}).All(&todos)
	if err != nil {
		return todos, err
	}
	return todos, nil
}