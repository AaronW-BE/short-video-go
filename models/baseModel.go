package models

import "log"

type BaseModel struct {
}

func (model *BaseModel) Find() {
	log.Println(*model, "find func")
}

func (model *BaseModel) FindAll() {
	log.Println(*model, "find all func")
}

func (model *BaseModel) FindOne() {
	log.Println(*model, "find one func")
}
