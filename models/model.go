package models

import "gopkg.in/mgo.v2/bson"

type BibleSentence struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	TodaysSentence string        `bson:"todays_sentence" json:"todays_sentence"`
	Date           string        `json:"date"`
}
