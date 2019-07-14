package repositories

import (
	. "ChristTheKing/models"
	"gopkg.in/mgo.v2/bson"
)

type BibleRepository struct{}

func (*BibleRepository) GetTodaysQuote() (bs BibleSentence, err error) {
	bs = BibleSentence{}
	sessionCopy := DatabaseSession.Copy()
	c := sessionCopy.DB(CTK_DATABASE).C(COL_BIBLE_SENTENCE)
	// get the count of total documents in collection
	totalRecords, err := c.Count()
	// Fetch the latest inserted document
	err = c.Find(nil).Skip(totalRecords - 1).One(&bs)
	// Recover panic and return error
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
		sessionCopy.Close()
	}()
	return bs, err
}

func (*BibleRepository) AddTodaysQuote(bs BibleSentence) (err error) {
	sessionCopy := DatabaseSession.Copy()
	c := sessionCopy.DB(CTK_DATABASE).C(COL_BIBLE_SENTENCE)
	err = c.Insert(bs)
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
		sessionCopy.Close()
	}()
	return err
}

func (*BibleRepository) UpdateTodaysQuote(bs BibleSentence) (err error) {
	sessionCopy := DatabaseSession.Copy()
	c := sessionCopy.DB(CTK_DATABASE).C(COL_BIBLE_SENTENCE)
	err = c.UpdateId(bs.ID, bson.M{"$set": bson.M{"todays_sentence": bs.TodaysSentence,
		"date": bs.Date}})
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
		sessionCopy.Close()
	}()
	return err
}
