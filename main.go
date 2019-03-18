package main

import (
	"ChristTheKing/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Christ the king...")
	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
