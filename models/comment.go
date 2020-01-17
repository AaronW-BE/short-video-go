package models

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type Comment struct {
	Id      bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	VideoId bson.ObjectId `json:"video_id" bson:"video_id"`
	//TargetId bson.ObjectId `json:"target_id" bson:"target_id"` // 目标评论id
	OwnerId    bson.ObjectId `json:"owner_id" bson:"owner_id"` // 拥有者id
	Content    string        `json:"content" bson:"content"`
	CreateTime time.Time     `json:"create_time" bson:"create_time"`
	BaseModel  `json:"-" bson:"-"`
}

func (c *Comment) Create() (err error) {
	c.Name = "comments"
	log.Println("will create comment")

	c.CreateTime = time.Now()
	err = c.__create(c)
	return
}
