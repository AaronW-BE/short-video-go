package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func NewRandomNumber(len int) string {
	var str string
	for i := 0; i < len; i++ {
		rand.Seed(time.Now().Unix() + int64(i))
		str += strconv.Itoa(rand.Intn(10))
	}
	return str
}
