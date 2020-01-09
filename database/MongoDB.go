package database

import (
	"MiniVideo/utils"
	"time"
)
import (
	"gopkg.in/mgo.v2"
)

var Mongo *mgo.Session

func init()  {
	var err error

	config := utils.LoadConfig()

	dialInfo := &mgo.DialInfo{
		Addrs:          []string{config.Mongodb.Host},
		Username:       config.Mongodb.User,
		Password:       config.Mongodb.Password,
		Timeout:		time.Second * time.Duration(config.Mongodb.Timeout),
	}
	Mongo, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		panic("db init failed")
	}
}
