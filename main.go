package main

import (
	"MiniVideo/database"
	"MiniVideo/router"
)

func main() {
	router.InitRoute()
	defer database.Mongo.Close()

	//pwd, _ := utils.EncryptStringToPassword("123456")
	//u := &models.User{
	//	Username: "admin",
	//	Password: pwd,
	//}
	//u.Create()

}
