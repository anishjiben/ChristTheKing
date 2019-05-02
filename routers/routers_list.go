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
	Route{"GetDailyBibleSentence",
		"GET",
		"/get_bible_quote",
		handlers.GetDailyBibleQote},
	Route{"PostDailyBibleSentence",
		"POST",
		"/post_bible_quote",
		handlers.PostDailyBibleQuote},
	Route{"AddUser",
		"POST",
		"/sign_up",
		handlers.SignUpUser},
	Route{"Login",
		"POST",
		"/login",
		handlers.UserLoginIn},
}
