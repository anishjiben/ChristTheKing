package handlers

import (
	. "ChristTheKing/errors"
	"ChristTheKing/repositories"
	"encoding/json"
	"net/http"
)

// Handler to serve the daily bible quote
func DailyBibleSentence(w http.ResponseWriter, r *http.Request) {

	// Get daily quote from the bible_repo repository
	bibleSentence, err := repositories.GetTodaysQuote()
	w.Header().Set("Content-Type", "applicatio/json:charset=UTF-8")

	if err != nil {
		// Send the error as response, if data fetch from database fails
		w.WriteHeader(http.StatusInternalServerError)
		data, _ := json.Marshal(GetErrorMessage(CONNECTION_ERROR))
		w.Write(data)
	} else {
		// Send the daily bible sentence as response, if data fetch from database is success
		w.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(bibleSentence)
		w.Write(data)
	}
}
