package repositories

import (
	. "ChristTheKing/models"
	"golang.org/x/crypto/bcrypt"
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

func (*UserRepository) IsAuthorizedUser(credential UserCredential) (authorizedUser bool, err error) {
	user := CtkUser{}
	sessionCopy := DatabaseSession.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(CTK_DATABASE).C(COL_USERS)
	// Get the user from db whose match with the userName
	if err = c.Find(bson.M{"user_name": credential.UserName}).One(&user); err != nil {
		return false, err
	}
	// Compare the password from database and received password.
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credential.Password)); err != nil {
		// Password doesnot matches hence return with error
		return false, err
	}
	// Password matches, hence return with true.
	return true, err
}
