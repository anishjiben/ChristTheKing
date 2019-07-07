package repositories

import (
	. "ChristTheKing/models"
	"gopkg.in/mgo.v2/bson"
)

type TokenRepository struct{}

func (*TokenRepository) SaveToken(jwtToken string) (err error) {
	sessionCopy := DatabaseSession.Copy()
	c := sessionCopy.DB(CTK_DATABASE).C(COL_TOKEN)
	// Token struct
	token := Token{JwtToken: jwtToken}
	err = c.Insert(token)
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
		sessionCopy.Close()
	}()
	return err
}

func (*TokenRepository) IsTokenExist(jwtToken string) (token Token, err error) {
	sessionCopy := DatabaseSession.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(CTK_DATABASE).C(COL_TOKEN)
	err = c.Find(bson.M{"token": jwtToken}).One(&token)
	return token, err
}

func (*TokenRepository) DeleteBlacklistedTokens() error {
	sessionCopy := DatabaseSession.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(CTK_DATABASE).C(COL_TOKEN)
	_, err := c.RemoveAll(nil)
	return err
}
