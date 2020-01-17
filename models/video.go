package models

import (
	types2 "MiniVideo/types"
	"go/types"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Video struct {
	Id           bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	OwnerId      bson.ObjectId `json:"owner_id" bson:"owner_id"` // 拥有者id
	Title        string        `json:"title" form:"title" bson:"title"`
	Description  string        `json:"description" form:"description" bson:"description"`
	Thumb        string        `json:"thumb" form:"thumb" bson:"thumb"`
	Likes        int           `json:"likes" bson:"likes"`
	CommentCount int           `json:"comment_count" bson:"comment_count"`
	ShareCount   int           `json:"share_count" bson:"share_count"`
	IsHidden     bool          `json:"is_hidden" bson:"is_hidden"`
	IsIllegal    bool          `json:"is_illegal" bson:"is_illegal"`
	InterTags    types.Array   `json:"-" bson:"inter_tags"`
	CreateTime   time.Time     `json:"create_time" bson:"create_time"`
	BaseModel    `bson:"-"`
}

func (v *Video) FindVideosWithPagination(selector interface{}, page int, size int, sort bson.M) (types2.PaginateResult, error) {
	v.Name = "videos"
	return v.findWithPagination(selector, page, size, sort, &[]Video{})
}
