package models

import (
	"MiniVideo/database"
	"gopkg.in/mgo.v2"
	"log"
)

type BaseModel struct {
	Name string
}

func (model *BaseModel) getCollection() *mgo.Collection {
	return database.Mongo.DB("douyin").C(model.Name)
}

func (m *BaseModel) checkName() {
	if m.Name == "" {
		panic("base model needs name filed")
	}
}

func (model *BaseModel) find() {
	model.checkName()

	log.Println(*model, "find func")
}

func (model *BaseModel) findAll(docList interface{}) error {
	model.checkName()
	log.Println(*model, "find all func")
	return model.getCollection().Find(nil).All(docList)
}

func (model *BaseModel) findOne() {
	model.checkName()
	log.Println(*model, "find one func")
}

func (model *BaseModel) create(m interface{}) error {
	model.checkName()
	return model.getCollection().Insert(m)
}

//========================== update ===============================

func (model *BaseModel) update(selector interface{}, update interface{}) error {
	model.checkName()
	return model.getCollection().Update(selector, update)
}

func (model *BaseModel) updateId(id interface{}, update interface{}) error {
	model.checkName()
	return model.getCollection().Update(id, update)
}

func (model *BaseModel) updateAll(selector interface{}, update interface{}) error {
	model.checkName()
	return model.getCollection().Update(selector, update)
}

//========================== delete ===============================

func (model *BaseModel) delete(selector interface{}) error {
	model.checkName()
	return model.getCollection().Remove(selector)
}

func (model *BaseModel) RemoveId(id interface{}) error {
	model.checkName()
	return model.getCollection().RemoveId(id)
}

func (model *BaseModel) RemoveAll(selector interface{}) (*mgo.ChangeInfo, error) {
	model.checkName()
	return model.getCollection().RemoveAll(selector)
}
