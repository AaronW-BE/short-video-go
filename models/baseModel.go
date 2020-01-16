package models

import (
	"MiniVideo/database"
	"MiniVideo/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type BaseModel struct {
	Name string
}

func (model *BaseModel) getCollection() *mgo.Collection {
	model.checkName()
	return database.Mongo.DB("douyin").C(model.Name)
}

func (m *BaseModel) checkName() {
	if m.Name == "" {
		panic("base model needs name filed")
	}
}

func (model *BaseModel) find() {
	log.Println(*model, "find func")
}

func (model *BaseModel) findAll(docList interface{}) error {
	return model.getCollection().Find(nil).All(docList)
}

func (model *BaseModel) findOne(selector interface{}) (doc interface{}, err error) {
	err = model.getCollection().Find(selector).One(&doc)
	return
}

func (model *BaseModel) findWithPagination(selector interface{}, page int, size int, sort bson.M) (result types.PaginateResult, err error) {
	pipeM := []bson.M{
		{"$match": selector},
		{"$skip": (page - 1) * size},
		{"$limit": size},
		{"$sort": sort},
	}
	pipe := model.getCollection().Pipe(pipeM)
	result = types.PaginateResult{
		List:  nil,
		Page:  page,
		Size:  size,
		Total: 0,
	}

	// TODO 优化返回列表数据类型
	//

	err = pipe.All(&result.List)
	count, err := model.getCollection().Find(selector).Count()
	result.Total = count
	return
}

func (model *BaseModel) create(m interface{}) error {
	return model.getCollection().Insert(m)
}

//========================== update ===============================

func (model *BaseModel) update(selector interface{}, update interface{}) error {
	return model.getCollection().Update(selector, update)
}

func (model *BaseModel) updateId(id interface{}, update interface{}) error {
	return model.getCollection().Update(id, update)
}

func (model *BaseModel) updateAll(selector interface{}, update interface{}) error {
	return model.getCollection().Update(selector, update)
}

//========================== delete ===============================

func (model *BaseModel) delete(selector interface{}) error {
	return model.getCollection().Remove(selector)
}

func (model *BaseModel) RemoveId(id interface{}) error {
	return model.getCollection().RemoveId(id)
}

func (model *BaseModel) RemoveAll(selector interface{}) (*mgo.ChangeInfo, error) {
	return model.getCollection().RemoveAll(selector)
}
