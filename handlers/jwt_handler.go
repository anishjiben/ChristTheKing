package handlers

import "time"
import "github.com/dgrijalva/jwt-go"

type CtkClaims struct {
	userName       string
	standardClaims jwt.StandardClaims
}
type JWTAuthentication struct {
	signatureKey []byte
}

var jwtAuthInstance *JWTAuthentication

func InitializeJWTAuthentication() *JWTAuthentication {
	if jwtAuthInstance == nil {
		// TODO: Get the secret_key from the environment variable, dont encode it here
		jwtAuthInstance = &JWTAuthentication{[]byte("secret_key")}
	}
	return jwtAuthInstance
}

func (jwtInstance *JWTAuthentication) GenerateToken(userName string) (token string, err error) {
	// Expiration time of token, 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
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
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (i interface{}, e error) {
		return jwtAuthInstance.signatureKey, nil
	})
	return token.Valid, err
}
