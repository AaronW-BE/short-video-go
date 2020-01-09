package models

import "MiniVideo/database"

type Video struct {
	Id int
}

func (video *Video) Insert() (id int, err error)  {
	videos := database.Mongo.DB("douyin").C("videos")
	id = video.Id
	err = videos.Insert(video)
	return
}