package handlers

import (
	. "ChristTheKing/messages"
	. "ChristTheKing/models"
	"ChristTheKing/repositories"
	. "ChristTheKing/validators"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

var userRepository repositories.UserRepository

func init() {
	userRepository = repositories.UserRepository{}
}

func SignUpUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 8)
	if err != nil {
		panic(err)
		return
	}
	user := CtkUser{"", r.FormValue("user_name"), hashedPwd, r.FormValue("role"), r.FormValue("account_created_by"), time.Now().UTC()}
	// validate if the form data is valid
	if err := Validate.Struct(user); err != nil {
		sendResponse(w, http.StatusBadRequest, GetErrorMessage(BAD_REQUEST))
		return
	}

	// Check if user name already exist
	if isExist, _ := userRepository.IsUserExist(user.UserName); isExist {
		sendResponse(w, http.StatusConflict, GetErrorMessage(USER_EXIST))
		return
	} else if err := userRepository.AddUser(user); err != nil {
		sendResponse(w, http.StatusInternalServerError, GetErrorMessage(INSERTION_FAILED))
	} else {
		sendResponse(w, http.StatusCreated, GetSuccessMessage(INSERTED_SUCCESSFULLY))
	}

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			sendResponse(w, http.StatusInternalServerError, err)
		}
	}()
}

func UserLoginIn(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userCredential := UserCredential{r.FormValue("user_name"), r.FormValue("password")}
	// Validate if the form data is valid
	if err := Validate.Struct(userCredential); err != nil {
		sendResponse(w, http.StatusBadRequest, GetErrorMessage(BAD_REQUEST))
		return
	}
	// Check for authorization
	if authorizedUser, _ := userRepository.IsAuthorizedUser(userCredential); !authorizedUser {
		// Un Authorized
		sendResponse(w, http.StatusUnauthorized, GetErrorMessage(UN_AUTHORIZED_USER))
		return
	}
	token, err := jwtAuthInstance.GenerateToken(userCredential.UserName)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, GetErrorMessage(TOKEN_CREATION_FAILED))
		return
	}
	// Adding the jwt-token to the response header
	w.Header().Add("Authorization", "Bearer "+token)
	// Authorized user, hence logged-in succesfully
	sendResponse(w, http.StatusOK, GetSuccessMessage(LOGGEDIN_SUCCESFULLY))
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	knownUser := false
	var err error
	r.ParseForm()
	// Token Validation
	jwtToken := r.Header.Get("Authorization")
	splitToken := strings.Split(jwtToken, "Bearer ")
	if len(splitToken) > 1 {
		knownUser, err = jwtAuthInstance.VerifyToken(splitToken[1])
	}
	if !knownUser || err != nil {
		sendResponse(w, http.StatusUnauthorized, GetErrorMessage(LOGIN_REQUIRED))
		return
	}
	// Black list the token
	if blackListed := jwtAuthInstance.BlackListToken(splitToken[1]); !blackListed {
		err = errors.New("black listing token unsuccesfull")
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}
	//On Logged out succesfully
	sendResponse(w, http.StatusOK, GetSuccessMessage(LOGGEDOUT_SUCCESFULLY))

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			sendResponse(w, http.StatusInternalServerError, err)
		}
	}()
}
