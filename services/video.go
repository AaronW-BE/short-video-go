package services

import "MiniVideo/models"

type Video struct {
	Id int `json:"id"`
}

func (video *Video)Insert() (id int, err error)  {
	var model = models.Video{}
	model.Id = video.Id

	id, err = model.Insert()
	return
}
