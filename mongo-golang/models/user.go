package models

import ("gopkg.in/mgo.v2/bson"
// "time"
)

type User struct{
	Id    bson.ObjectId    `json:"id"  bson:"_id"`
	Name  string           `json:"name" bson:"name"`
	Email string          `json:"email" bson:"email"`
	Password   string      `json:"password" bson:"password"`
}

type Post struct{
	Id    string    `json:"id1"  bson:"_id1"`
	Caption  string           `json:"caption" bson:"caption"`
	ImageURL string          `json:"image" bson:"image"`
	// Password   string      `json:"password" bson:"password"`
	PostTimestamp bson.MongoTimestamp   `json:"time" bson:"time"`
}