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
