package models

import (
	types2 "MiniVideo/types"
	"go/types"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Video struct {
	Id           bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	OwnerId      bson.ObjectId `bson:"owner_id"` // 拥有者id
	Title        string        `json:"title" form:"title" bson:"title"`
	Description  string        `json:"description" form:"description" bson:"description"`
	Thumb        string        `json:"thumb" form:"thumb" bson:"thumb"`
	Likes        int           `bson:"likes"`
	CommentCount int           `bson:"comment_count"`
	ShareCount   int           `bson:"share_count"`
	IsHidden     bool          `bson:"is_hidden"`
	IsIllegal    bool          `bson:"is_illegal"`
	InterTags    types.Array   `bson:"inter_tags"`
	CreateTime   time.Time     `bson:"create_time"`
	BaseModel
}

func (v *Video) FindVideosWithPagination(selector interface{}, page int, size int, sort bson.M) (types2.PaginateResult, error) {
	v.Name = "videos"
	return v.findWithPagination(selector, page, size, sort)
}
