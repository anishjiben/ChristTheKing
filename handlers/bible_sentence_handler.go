package handlers

import (
	"ChristTheKing/models"
	"encoding/json"
	"net/http"
)

func DailyBibleSentence(w http.ResponseWriter, r *http.Request) {
	bibleSentence := models.BibleSentence{
		TodaysQuote: "This is daily bible sentence (Mathew 21 :2-1)",
		Date:        "11/12/1234",
	}

	w.Header().Set("Content-Type", "applicatio/json:charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(bibleSentence)
	w.Write(data)
}
