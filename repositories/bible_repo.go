package repositories

import (
	. "ChristTheKing/models"
	"gopkg.in/mgo.v2"
)

func GetTodaysQuote() (bs BibleSentence, err error) {
	bs = BibleSentence{}
	session, err := mgo.DialWithInfo(MongoDBDialInfo)
	defer session.Close()
	if err != nil {
		return bs, err
	}
	c := session.DB(CTK_DATABASE).C(COL_BIBLE_SENTENCE)
	totalRecords, err := c.Count()
	err = c.Find(nil).Skip(totalRecords - 1).One(&bs)
	return bs, err
}
