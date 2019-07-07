package handlers

import (
	"ChristTheKing/repositories"
	"errors"
	"fmt"
	"github.com/robfig/cron"
	"time"
)
import "github.com/dgrijalva/jwt-go"

type CtkClaims struct {
	userName       string
	standardClaims jwt.StandardClaims
}
type JWTAuthentication struct {
	signatureKey    []byte
	tokenExpireTime time.Duration //In minutes
}

var jwtAuthInstance *JWTAuthentication
var tokenRepo repositories.TokenRepository
var cronShedule *cron.Cron

func InitializeJWTAuthentication() *JWTAuthentication {
	if jwtAuthInstance == nil {
		// TODO: Get the secret_key from the environment variable, dont hardcode it here
		jwtAuthInstance = &JWTAuthentication{[]byte("secret_key"), 2}
		tokenRepo = repositories.TokenRepository{}
	}
	return jwtAuthInstance
}

// Sheduler to clean the blacklisted tokens from the mongodb
func SheduleBlacklistTokensCleanJob() *cron.Cron {
	if cronShedule == nil {
		cronShedule = cron.New()
		cronShedule.AddFunc("@monthly", jwtAuthInstance.emptyBlackListedTokens)
		cronShedule.Start()
	}
	return cronShedule
}

func StopSheduledTokenCleanUpJob() {
	if cronShedule != nil {
		cronShedule.Stop()
	}
}

func (jwtInstance *JWTAuthentication) GenerateToken(userName string) (token string, err error) {
	// Expiration time of token, 5 minutes
	expirationTime := time.Now().Add(jwtInstance.tokenExpireTime * time.Minute)
	// Claims to be used while creating token
	ctkClaims := jwt.MapClaims{
		"userName": userName,
		"exp":      expirationTime.Unix(),
		"standardClaim": jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "CTK",
			IssuedAt:  time.Now().Unix(),
		},
	}
	// Generating the token with secret key
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS512, ctkClaims)
	token, err = unsignedToken.SignedString(jwtInstance.signatureKey)
	return token, err
}

func (jwtInstance *JWTAuthentication) VerifyToken(jwtToken string) (valid bool, err error) {
	// Check if token is valid
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (i interface{}, e error) {
		return jwtAuthInstance.signatureKey, nil
	})
	// Check if token is blacklisted
	blackListed := jwtInstance.IsTokenBlackListed(jwtToken)
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	return token.Valid && !blackListed, err
}

func (jwtInstance *JWTAuthentication) RefreshToken(jwtToken string) (renewdToken string, err error) {
	ctkClaims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(jwtToken, ctkClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtInstance.signatureKey, nil
	})
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	/* Check if the current time is about to(should not be 30sec befor expire time)
	expire or not, panic if the token is not expired*/
	expiredTime := int64(ctkClaims["exp"].(float64))
	if time.Unix(expiredTime, 0).Sub(time.Now()) > 30*time.Second {
		panic(errors.New("Time not expired"))
	}

	if blackListed := jwtInstance.BlackListToken(jwtToken); !blackListed {
		err = errors.New("black listing token unsuccesfull")
	}

	// Extend the time of token to 5 more minutes
	ctkClaims["exp"] = time.Now().Add(jwtInstance.tokenExpireTime * time.Minute).Unix()
	// Generating the token with secret key
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS512, ctkClaims)
	renewdToken, err = unsignedToken.SignedString(jwtInstance.signatureKey)
	return renewdToken, err
}

func (jwtInstance *JWTAuthentication) BlackListToken(jwtToken string) bool {
	if err := tokenRepo.SaveToken(jwtToken); err != nil {
		return false
	}
	return true
}

func (jwtInstance *JWTAuthentication) IsTokenBlackListed(jwtToken string) bool {
	if token, _ := tokenRepo.IsTokenExist(jwtToken); token.JwtToken == jwtToken {
		return true
	}
	return false
}

func (jwtInstance *JWTAuthentication) emptyBlackListedTokens() {
	if err := tokenRepo.DeleteBlacklistedTokens(); err != nil {
		//TODO: Use logger here instead of println
		fmt.Println("Cleaning unsessfull : ", err)
	}
}
