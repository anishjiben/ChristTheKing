package handlers

import (
	. "ChristTheKing/messages"
	. "ChristTheKing/models"
	"ChristTheKing/repositories"
	. "ChristTheKing/validators"
	"net/http"
	"strings"
)

var upcomingEventRepo repositories.UpcomingEventRepo

func init() {
	upcomingEventRepo = repositories.UpcomingEventRepo{}
}

func PostUpcomingEvents(w http.ResponseWriter, r *http.Request) {
	knownUser := false
	var err error
	r.ParseForm()
	ue := UpcomingEvent{
		"",
		r.FormValue("title"),
		r.FormValue("description"),
		r.FormValue("time"),
		r.FormValue("imageUrl")}
	// validate if the form data is valid
	if err := Validate.Struct(ue); err != nil {
		sendResponse(w, http.StatusBadRequest, GetErrorMessage(BAD_REQUEST))
		return
	}
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
	if err := upcomingEventRepo.SaveUpcomingEvent(ue); err != nil {
		// Send the error as response, if data insertion fails
		sendResponse(w, http.StatusInternalServerError, GetErrorMessage(CONNECTION_ERROR))
	} else {
		// Send the success message as response, if data has been inserted succesfully
		sendResponse(w, http.StatusCreated, GetSuccessMessage(INSERTED_SUCCESSFULLY))
	}
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			sendResponse(w, http.StatusInternalServerError, err)
		}
	}()
}

func GetUpcomingEvents(w http.ResponseWriter, r *http.Request) {
	// Get all upcoming events
	upcomingEvents, err := upcomingEventRepo.FetchUpcomingEvents()
	if err != nil {
		// Send the error as response, if data fetch from database fails
		sendResponse(w, http.StatusInternalServerError, GetErrorMessage(CONNECTION_ERROR))
	} else {
		// Send the upcoming events as response, if data fetch from database is success
		sendResponse(w, http.StatusOK, upcomingEvents)
	}
}
