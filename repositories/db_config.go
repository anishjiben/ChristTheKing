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

// Name of the collections in database "christ_the_king"
const COL_BIBLE_SENTENCE = "bible_sentences"
const COL_USERS = "users"
const COL_TOKEN = "blacklisted_tokens"
const COL_UPCOMING_EVENTS = "upcoming_events"

func init() {
	var err error
	if DatabaseSession, err = mgo.DialWithInfo(MongoDBDialInfo); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Connected..")
}
