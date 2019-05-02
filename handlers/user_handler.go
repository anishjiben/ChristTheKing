package handlers

import (
	. "ChristTheKing/messages"
	. "ChristTheKing/models"
	"ChristTheKing/repositories"
	. "ChristTheKing/validators"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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
		sendResponse(w, http.StatusUnauthorized, UN_AUTHORIZED_USER)
		return
	}
	// Authorized user, hence logged-in succesfully
	sendResponse(w, http.StatusOK, LOGGEDIN_SUCCESFULLY)
}
