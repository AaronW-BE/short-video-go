package models

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

const LikeType_Video int = 1
const LikeType_Comment int = 2

type Like struct {
	Id         bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Type       int           `json:"type" bson:"type"`
	Uid        bson.ObjectId `json:"uid" bson:"user_id"`
	Vid        bson.ObjectId `json:"vid" bson:"video_id"`
	CreateTime time.Time     `json:"create_time" bson:"create_time"`
	BaseModel  `bson:"-"`
}

// 创建点赞
func NewLike(likeData *Like) error {
	like := &Like{}
	video := &Video{}

	like.Name = "likes"
	video.Name = "videos"

	foundLike := Like{}
	likeFounded := like.Find(bson.M{
		"user_id":  likeData.Uid,
		"video_id": likeData.Vid,
	}).One(&foundLike)

	foundVideo := Video{}
	videoFindErr := video.FindById(likeData.Vid).One(&foundVideo)
	if videoFindErr != nil {
		return videoFindErr
	}

	if likeFounded == nil {
		// 删除点赞记录
		removeLikeErr := like.RemoveId(foundLike.Id)
		if removeLikeErr != nil {
			log.Println(removeLikeErr)
			return removeLikeErr
		}

		// 已点赞，取消点赞
		updateErr := video.update(bson.M{
			"_id": foundVideo.Id,
		}, bson.M{
			"likes": foundVideo.Likes - 1,
		})
		return updateErr
	} else {
		log.Println(likeData)
		likeData.CreateTime = time.Now()

		// 创建点赞记录
		createErr := like.__create(likeData)
		if createErr != nil {
			log.Println("create like err, ", createErr)
			return createErr
		}

		updateErr := video.update(bson.M{
			"_id": foundVideo.Id,
		}, bson.M{
			"likes": foundVideo.Likes + 1,
		})
		if updateErr != nil {
			log.Println("update err", updateErr)
			return updateErr
		}
	}

	return nil
}

func RemoveLike(vid bson.ObjectId) {
	// 查询视频，视频点赞减一

}
