package repository

import (
	"todoapp/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TodoRepository interface {
	CreateTodo(todo model.Todo) error
	DropTodo(todoId string) error
	GetTodos() ([]TodoDao, error)
}

type todoRepository struct {
	provider *mgo.Session
	dbName   string
}

func NewTodoRepository(provider *mgo.Session, dbName string) *todoRepository {
	return &todoRepository{
		provider: provider,
		dbName:   dbName,
	}
}

func (this *todoRepository) CreateTodo(todo model.Todo) error {
	session := this.provider.Copy()
	defer session.Close()
	todoCollection := session.DB(this.dbName).C("todo")
	objectId := bson.ObjectIdHex(todo.Id)
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

func (this *todoRepository) DropTodo(referenceId string) error {
	session := this.provider.Copy()
	defer session.Close()
	todoId := bson.ObjectIdHex(referenceId)
	err := session.DB(this.dbName).C("todo").RemoveId(todoId)
	return err
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
