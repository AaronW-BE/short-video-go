package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	STATE_NORMAL      = 1  // 正常
	STATE_BLOCKED     = 0  // 禁止登录
	STATE_NEED_VERIFY = -1 // 需要验证
)

type User struct {
	Id        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string        `bson:"username"`
	Password  string        `bson:"password"`
	Phone     string        `bson:"phone"`
	State     int           `bson:"state"`
	Register  time.Time     `bson:"register"`
	LastLogin time.Time     `bson:"last_login"`
	BaseModel `bson:"-"`
}

func (u *User) FindAll() (users []User) {
	u.Name = "users"
	_ = u.findAll(&users)
	return
}

func (u *User) FindOne(selector interface{}) (user User, err error) {
	u.Name = "users"
	user, err = u.findOne(selector)
	return
}

func (u *User) Create() (err error) {
	u.Name = "users"
	u.State = STATE_NORMAL
	u.Register = time.Now()
	err = u.create(u)
	return
}
