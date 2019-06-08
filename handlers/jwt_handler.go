package handlers

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

func (jwtInstance *JWTAuthentication) GenerateToke() {

}
