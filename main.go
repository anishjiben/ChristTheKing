package main

import (
	"ChristTheKing/handlers"
	"ChristTheKing/routers"
	"fmt"
	"log"
	"net/http"
)

func init() {
	handlers.InitializeJWTAuthentication()
	handlers.SheduleBlacklistTokensCleanJob()
}
func main() {
	fmt.Println("Christ the king... service started")
	router := routers.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(handlers.UrlMatchNotFoundHandler)
	log.Fatal(http.ListenAndServe(":8000", router))

	defer func() {
		handlers.StopSheduledTokenCleanUpJob()
	}()
}
