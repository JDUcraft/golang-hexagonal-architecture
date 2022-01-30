package api

import (
	"fmt"
	"log"
	"main/homepage/application"
	"main/homepage/infrastructure/repository"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/home", HomePage)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	clock := repository.Clock{}
	_, err := fmt.Fprintf(w, application.ReadTime(clock))
	if err != nil {
		return
	}
}
