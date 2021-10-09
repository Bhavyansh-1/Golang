package models

import "gopkg.in/mgo.v2/bson"

// all the fields that were asked us
type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Password string        `json:"password" bson:"password"`
	Email    string        `json:"email" bson:"email"`
}

//here whenever you create a post you have to give the user id i have taken user id here as string any data type can be taken for it
type Post struct {
	Id       bson.ObjectId       `json:"id" bson:"_id"`
	Caption  string              `json:"caption" bson:"caption"`
	ImageUrl string              `json:"image" bson:"image"`
	Time     bson.MongoTimestamp `json:"time" bson:"time"`
	Userid   string              `json:"userid" bson:"userid"`
}
