package repositories

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

const CTK_SERVER = "localhost:27017"
const CTK_DATABASE = "christ_the_king"

var MongoDBDialInfo = &mgo.DialInfo{
	Addrs:    []string{CTK_SERVER},
	Timeout:  time.Duration(2e+10), // 20seconds
	Database: CTK_DATABASE,
	FailFast: true,
}

// Data base session
var DatabaseSession *mgo.Session

// Name of the collections in collection "christ_the_king"
const COL_BIBLE_SENTENCE = "bible_sentences"
const COL_USERS = "users"

func init() {
	var err error
	if DatabaseSession, err = mgo.DialWithInfo(MongoDBDialInfo); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Connected..")
}
