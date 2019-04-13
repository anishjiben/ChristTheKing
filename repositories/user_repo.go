package repositories

import (
	. "ChristTheKing/models"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct{}

func (*UserRepository) IsUserExist(userName string) (isExist bool, err error) {
	sessionCopy := DatabaseSession.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(CTK_DATABASE).C(COL_USERS)
	count, err := c.Find(bson.M{"user_name": userName}).Count()
	return count > 0, err
}

func (*UserRepository) AddUser(user CtkUser) (err error) {
	sessionCopy := DatabaseSession.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(CTK_DATABASE).C(COL_USERS)
	err = c.Insert(user)
	return err
}
