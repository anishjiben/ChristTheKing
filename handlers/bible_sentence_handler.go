package handlers

import (
	"ChristTheKing/repositories"
	"encoding/json"
	"net/http"
)

func DailyBibleSentence(w http.ResponseWriter, r *http.Request) {

	bibleSentence, err := repositories.GetTodaysQuote()
	w.Header().Set("Content-Type", "applicatio/json:charset=UTF-8")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data, _ := json.Marshal(repositories.CONNECTION_ERROR)
		w.Write(data)
	} else {
		w.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(bibleSentence)
		w.Write(data)
	}
}
