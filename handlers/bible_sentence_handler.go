package handlers

import (
	. "ChristTheKing/messages"
	. "ChristTheKing/models"
	"ChristTheKing/repositories"
	. "ChristTheKing/validators"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

var bibleRepo repositories.BibleRepository

func init() {
	bibleRepo = repositories.BibleRepository{}
}

// Handler to serve daily bible quote
func GetDailyBibleQote(w http.ResponseWriter, r *http.Request) {
	// Get daily quote from the bible_repo repository
	bibleSentence, err := bibleRepo.GetTodaysQuote()
	if err != nil {
		// Send the error as response, if data fetch from database fails
		sendResponse(w, http.StatusInternalServerError, GetErrorMessage(CONNECTION_ERROR))
	} else {
		// Send the daily bible sentence as response, if data fetch from database is success
		sendResponse(w, http.StatusOK, bibleSentence)
	}
}

// Handler to post daily bible quote
func PostDailyBibleQuote(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	bs := BibleSentence{"", r.FormValue("todays_sentence"), time.Now().UTC()}
	// validate if the form data is valid
	if err := Validate.Struct(bs); err != nil {
		sendResponse(w, http.StatusBadRequest, GetErrorMessage(BAD_REQUEST))
		return
	}
	// Token Validation
	jwtToken := r.Header.Get("Authorization")
	splitToken := strings.Split(jwtToken, "Bearer ")
	knownUser, err := jwtAuthInstance.VerifyToken(splitToken[1])
	if !knownUser || err != nil {
		sendResponse(w, http.StatusUnauthorized, GetErrorMessage(LOGIN_REQUIRED))
		return
	}
	if err := bibleRepo.AddTodaysQuote(bs); err != nil {
		// Send the error as response, if data inserts fails
		sendResponse(w, http.StatusInternalServerError, GetErrorMessage(CONNECTION_ERROR))
	} else {
		// Send the succes message as response, if data has been isert succesfully
		sendResponse(w, http.StatusCreated, GetSuccessMessage(INSERTED_SUCCESSFULLY))
	}
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			sendResponse(w, http.StatusInternalServerError, err)
		}
	}()
}

// Writes the response to the ResponseWriter
func sendResponse(w http.ResponseWriter, statusCode int, responsBody interface{}) {
	w.Header().Set("Content-Type", "applicatio/json:charset=UTF-8")
	w.WriteHeader(statusCode)
	data, _ := json.Marshal(responsBody)
	w.Write(data)
}
