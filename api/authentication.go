package api

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	//
	//TODO: @omid do this part
	_, err := fmt.Fprint(w, "login done")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(200)
}


func SignUp(w http.ResponseWriter, r *http.Request) {
	//
	//TODO: @omid do this part

	_, err := fmt.Fprint(w, "signUp done")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(200)
}