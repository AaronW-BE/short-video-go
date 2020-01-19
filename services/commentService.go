package services

import (
	"MiniVideo/models"
	"MiniVideo/types"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

func FindCommentsByVideoId(id interface{}, page int, size int) (types.PaginateResult, error) {
	comment := models.Comment{}
	comment.Name = "videos"

	return comment.FindWithPagination(bson.M{
		"video_id": id,
	}, page, size, bson.M{
		"create_time": -1,
	}, []models.Comment{})
}

func PublishVideoComment(id string, content string, userId bson.ObjectId) (err error) {
	comment := models.Comment{}
	comment.Name = "videos"

	// 检查视频是否存在
	exist := VideoExist(bson.ObjectIdHex(id))
	if !exist {
		return errors.New("视频不存在")
	}

	comment.VideoId = bson.ObjectIdHex(id)
	comment.Content = content
	comment.OwnerId = userId
	//comment.TargetId = ""

	err = comment.Create()

	// 更新视频评论数

	return err
}

func LikeVideo(uid bson.ObjectId, vid bson.ObjectId) error {
	like := &models.Like{
		Type: models.LikeType_Video,
		Uid:  uid,
		Vid:  vid,
	}
	return models.NewLike(like)
}
