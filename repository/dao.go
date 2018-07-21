package repository

import "gopkg.in/mgo.v2/bson"

type TodoDao struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	// TodoItems []TodoItemDao `bson:"items"`
	IsPending bool `bson:"isPending"`
}

// type TodoItemDao struct {
// 	Id          int `bson:"_id"`
// 	Description int `bson:"description"`
// }
