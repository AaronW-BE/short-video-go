package models

import (
	"MiniVideo/database"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	CollectionName string = "videos"
	Collection     *mgo.Collection
)

func init() {
	Collection = database.Mongo.DB("douyin").C(CollectionName)
}

type Video struct {
	Id          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string        `json:"title" form:"title" bson:"title"`
	Description string        `json:"description" form:"description" bson:"description"`
	Thumb       string        `json:"thumb" form:"thumb" bson:"thumb"`
	BaseModel
}

func (video *Video) Insert() (id int, err error) {
	err = Collection.Insert(video)
	return
}

func (video *Video) DeleteById() (err error) {
	err = Collection.RemoveId(video.Id)
	return
}

func findAll(query mgo.Query) (list []Video, err error) {
	err = Collection.Find(query).All(&list)
	return
}

func findOne(query mgo.Query) (video Video, err error) {
	err = Collection.Find(query).One(&video)
	return
}
