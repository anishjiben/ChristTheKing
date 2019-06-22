package handlers

import (
	. "ChristTheKing/messages"
	"net/http"
	"strings"
)

func RefreshUserToken(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	jwtToken := r.Header.Get("Authorization")
	splitToken := strings.Split(jwtToken, "Bearer ")

	// Form Validation
	if len(splitToken) < 2 {
		sendResponse(w, http.StatusBadRequest, GetErrorMessage(BAD_REQUEST))
		return
	}
	// Token validation
	knownUser, err := jwtAuthInstance.VerifyToken(splitToken[1])
	if !knownUser || err != nil {
		sendResponse(w, http.StatusUnauthorized, GetErrorMessage(LOGIN_REQUIRED))
		return
	}
	// Refreshing token
	renewdToken, err := jwtAuthInstance.RefreshToken(splitToken[1])
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, GetErrorMessage(TOKEN_CREATION_FAILED))
		return
	}
	// Adding the jwt-token to the response header
	w.Header().Add("Authorization", "Bearer "+renewdToken)
	// Authorized user, hence logged-in succesfully
	sendResponse(w, http.StatusOK, GetSuccessMessage(TOKEN_REFRESHED_SUCCESFULLY))
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			sendResponse(w, http.StatusInternalServerError, err)
		}
	}()
}
