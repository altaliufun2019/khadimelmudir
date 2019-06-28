package authentication

import (
	"context"
	"encoding/json"
	"fmt"
	"mudiralmaham/models"
	"mudiralmaham/utils/Logger"
	"mudiralmaham/utils/database"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	response := loginResponse{}
	credits, err := loginDecoder(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//TODO: check with database

	response.Msg = "login successful"
	response.Token, err = jwtEncoder(credits.Username)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	_, err = w.Write(output)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(200)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	response := signUpResponse{}
	// example for writing to db
	credits, err := signUpDecoder(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	_, err = database.DB.
		Collection("user").
		InsertOne(
			context.TODO(),
			models.User{
				Username: credits.Username,
				Name:     credits.Name,
				Password: credits.Password,
				Event:    models.Event{},
			})
	if err != nil {
		Logger.ErrorLogger.Println("Error in inserting file:", err)
		http.Error(w, err.Error(), 500)
		return
	}
	//TODO: check unique username
	response.Token, err = jwtEncoder(credits.Username)
	response.Msg = "signUp complete"
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = w.Write(output)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(200)
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "hello to you sir")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
}
