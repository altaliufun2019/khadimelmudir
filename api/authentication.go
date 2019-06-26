package api

import (
	"context"
	"fmt"
	"mudiralmaham/models"
	"mudiralmaham/utils/Database"
	"mudiralmaham/utils/Logger"
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
	// example for writing to db
	a := models.User{}
	a.Name = "abas"
	a.Password = "1234"
	a.Username = "somelkh"
	result, err := Database.Client.Database("makhzan").Collection("user").InsertOne(context.TODO(), a)
	if err != nil {
		Logger.ErrorLogger.Println("Error in inserting file:", err)
		return
	}
	Logger.GeneralLogger.Println(result.InsertedID)
	//
	//TODO: @omid do this part

	_, err = fmt.Fprint(w, "signUp done")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(200)
}