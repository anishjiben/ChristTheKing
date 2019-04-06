package repositories

import (
	. "ChristTheKing/models"
	"gopkg.in/mgo.v2"
)

func GetTodaysQuote() (bs BibleSentence, err error) {
	bs = BibleSentence{}
	session, err := mgo.DialWithInfo(MongoDBDialInfo)
	if err != nil {
		return bs, err
	}
	c := session.DB(CTK_DATABASE).C(COL_BIBLE_SENTENCE)
	// get the count of total documents in collection
	totalRecords, err := c.Count()
	// Fetch the latest inserted document
	err = c.Find(nil).Skip(totalRecords - 1).One(&bs)
	// Recover panic and return error
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
		session.Close()
	}()
	return bs, err
}
