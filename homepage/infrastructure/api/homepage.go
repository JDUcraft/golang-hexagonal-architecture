package api

import (
	"fmt"
	"main/homepage/application"
	"main/homepage/infrastructure/repository"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	clock := repository.Clock{}
	_, err := fmt.Fprintf(w, application.ReadTime(clock))
	if err != nil {
		return
	}
}
