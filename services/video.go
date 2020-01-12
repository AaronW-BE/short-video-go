package services

import (
	"MiniVideo/database"
	"MiniVideo/models"
)

type Video struct {
	Title       string `form:"title"`
	Description string `form:"description"`
}

func (video *Video) Insert() (id int, err error) {
	v := models.Video{
		Title:       video.Title,
		Description: video.Description,
	}
	id, err = v.Insert()
	return
}

func List() (videos []models.Video, err error) {
	videos = []models.Video{}
	video := models.Video{
		Id: "sdafds",
	}
	video.Find()
	err = database.Mongo.DB("douyin").C("videos").Find(nil).All(&videos)
	return
}
