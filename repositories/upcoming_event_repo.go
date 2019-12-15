package repositories

import (
	. "ChristTheKing/models"
	"gopkg.in/mgo.v2/bson"
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

func (*UpcomingEventRepo) UpdateUpcomingEvent(ue UpcomingEvent) (err error) {
	sessionCopy := DatabaseSession.Copy()
	c := sessionCopy.DB(CTK_DATABASE).C(COL_UPCOMING_EVENTS)
	err = c.UpdateId(ue.ID, bson.M{"$set": bson.M{"title": ue.Title,
		"description": ue.Description, "time": ue.Time, "image_url": ue.ImageUrl}})
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
		sessionCopy.Close()
	}()
	return err
}

func (*UpcomingEventRepo) RemoveUpcomingEvent(id string) (err error) {
	sessionCopy := DatabaseSession.Copy()
	c := sessionCopy.DB(CTK_DATABASE).C(COL_UPCOMING_EVENTS)
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			log.Println("panic occurred:", err)
		}
		sessionCopy.Close()
	}()
	err = c.RemoveId(bson.ObjectIdHex(id))

	return err
}
