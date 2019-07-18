package authentication

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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
	err = checkUserInDB(w, r, credits)
	if err != nil {
		return
	}

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

func checkUserInDB(w http.ResponseWriter, r *http.Request, credits loginCredentials) error {
	var user models.User
	err := database.
		DB.
		Collection("user").
		FindOne(context.TODO(), bson.D{{"username", credits.Username}}).
		Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return err
	}
	if user.Password != credits.Password {
		http.Error(w, "wrong password", 400)
		return errors.New("wrong password")
	}
	return nil
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	response := signUpResponse{}
	credits, err := signUpDecoder(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var user models.User
	err = database.
		DB.
		Collection("user").
		FindOne(context.TODO(), bson.D{{"username", credits.Username}}).
		Decode(&user)
	if user.Username != "" {
		http.Error(w, "duplicate username", 400)
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
				Projects: []models.Project{},
			})
	if err != nil {
		Logger.ErrorLogger.Println("Error in inserting file:", err)
		http.Error(w, err.Error(), 500)
		return
	}
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
