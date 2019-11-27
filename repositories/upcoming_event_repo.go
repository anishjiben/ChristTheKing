package repositories

import (
	. "ChristTheKing/models"
	"log"
)

type UpcomingEventRepo struct{}

// Handler to post Upcoming Event
func (*UpcomingEventRepo) SaveUpcomingEvent(ue UpcomingEvent) (err error) {
	sessionCopy := DatabaseSession.Copy()
	c := sessionCopy.DB(CTK_DATABASE).C(COL_UPCOMING_EVENTS)
	err = c.Insert(ue)
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
		sessionCopy.Close()
	}()
	return err
}

func (*UpcomingEventRepo) FetchUpcomingEvents() (ue []UpcomingEvent, err error) {
	sessionCopy := DatabaseSession.Copy()
	c := sessionCopy.DB(CTK_DATABASE).C(COL_UPCOMING_EVENTS)
	cursor := c.Find(nil)
	if err != nil {
		log.Fatal(err)
	}
	// get a list of all returned documents and print them out
	// see the mongo.Cursor documentation for more examples of using cursors
	//var results []bson.M
	if err = cursor.All(&ue); err != nil {
		return ue, err
	}

	// Recover panic and return error
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
		sessionCopy.Close()
	}()
	return ue, err
}
