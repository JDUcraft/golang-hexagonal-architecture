package api

import (
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/home", HomePage)
	log.Fatal(http.ListenAndServe(":80", nil))
}
