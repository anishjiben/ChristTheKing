package handlers

import (
	"fmt"
	"net/http"
)

func DailyBibleSentence(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This will returns you the daily bible sentence")
}
