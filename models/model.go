package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type BibleSentence struct {
	ID             bson.ObjectId `bson:"_id,omitempty" json:"id"`
	TodaysSentence string        `bson:"todays_sentence" json:"todays_sentence" validate:"required,empty"`
	Date           time.Time     `json:"date"`
}

type CtkUser struct {
	ID               bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UserName         string        `bson:"user_name" json:"user_name" validate:"required,empty"`
	Password         []byte        `bson:"password" json:"password" validate:"required,empty"`
	Role             string        `bson:"role" json:"role" validate:"required,empty"`
	AccountCreatedBy string        `bson:"account_created_by", json:"account_created_by" validate:"required,empty"`
	SignUpDate       time.Time     `bson:"sign_up_date" json:"sign_up_date" validate:"required,empty"`
}

type UserCredential struct {
	UserName string `bson:"user_name" json:"user_name" validate:"required,empty"`
	Password string `bson:"password" json:"password" validate:"required,empty"`
}

type Token struct {
	JwtToken string `bson:"token" json:"token"`
}

type UpcomingEvent struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title       string        `bson:"title" json:"title" validate:"required,empty"`
	Description string        `bson:"description" json:"description" validate:"required,empty"`
	Time        string        `bson:"time" json:"time" validate:"required,empty"`
	ImageUrl    string        `bson:"image_url" json:"image_url" validate:"required,empty"`
}
