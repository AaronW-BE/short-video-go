package services

import (
	"MiniVideo/models"
	"gopkg.in/mgo.v2/bson"
)

type Video struct {
	Title       string `form:"title"`
	Description string `form:"description"`
}

func (video *Video) Insert() (id int, err error) {

	return
}

func FetchRecommendVideos(uid bson.ObjectId) []models.Video {
	return nil
}

func FetchHotVideos() (interface{}, error) {
	v := models.Video{}
	result, err := v.FindVideosWithPagination(bson.M{}, 1, 5, bson.M{
		"create_time": -1,
		"like":        -1,
	})
	return result, err
}
