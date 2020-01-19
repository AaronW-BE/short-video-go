package models

import (
	"MiniVideo/database"
	"MiniVideo/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BaseModel struct {
	Name string `json:"-" bson:"-"`
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

func (model *BaseModel) Find(selector interface{}) *mgo.Query {
	return model.getCollection().Find(selector)
}

func (model *BaseModel) FindById(id interface{}) *mgo.Query {
	return model.getCollection().FindId(id)
}

func (model *BaseModel) findAll(docList interface{}) error {
	return model.getCollection().Find(nil).All(docList)
}

func (model *BaseModel) findOne(selector interface{}) (doc interface{}, err error) {
	err = model.getCollection().Find(selector).One(&doc)
	return
}

func (model *BaseModel) FindWithPagination(selector interface{}, page int, size int, sort bson.M, typedSlice interface{}) (result types.PaginateResult, err error) {
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
	err = pipe.All(typedSlice)
	result.List = typedSlice
	count, err := model.getCollection().Find(selector).Count()
	result.Total = count
	return
}

func (model *BaseModel) __create(m interface{}) error {
	return model.getCollection().Insert(m)
}

//========================== update ===============================

func (model *BaseModel) update(selector interface{}, update interface{}) error {
	return model.getCollection().Update(selector, update)
}

func (model *BaseModel) updateById(id interface{}, update interface{}) error {
	return model.getCollection().UpdateId(id, update)
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
