package main

import (
	"MiniVideo/database"
	"MiniVideo/router"
)

func main()  {
	router.InitRoute()
	defer database.Mongo.Close()
}
