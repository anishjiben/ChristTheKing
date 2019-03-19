package routers

import (
	"ChristTheKing/handlers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	{"GetDailyBibleSentence",
		"GET",
		"/get_bible_sentence",
		handlers.DailyBibleSentence},
}
