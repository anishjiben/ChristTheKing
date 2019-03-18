package routers

import (
	"ChristTheKing/handlers"
	"net/http"
)
import "github.com/gorilla/mux"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	{"GetDailyBibleSentence",
		"GET",
		"/get_bible_sentence",
		handlers.DailyBibleSentence},
}
