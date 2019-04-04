package handlers

import (
	"fmt"
	"net/http"
)

func UrlMatchNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 - Not Found")
}
