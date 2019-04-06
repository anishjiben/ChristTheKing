package repositories

import (
	"gopkg.in/mgo.v2"
)

const CTK_SERVER = "localhost:27019"
const CTK_DATABASE = "christ_the_king"

var MongoDBDialInfo = &mgo.DialInfo{
	Addrs:    []string{CTK_SERVER},
	Timeout:  600,
	Database: CTK_DATABASE,
	FailFast: true,
}

// Name of the collections in collection "christ_the_king"
const COL_BIBLE_SENTENCE = "bible_sentences"
