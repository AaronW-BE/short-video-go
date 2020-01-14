package services

import (
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
