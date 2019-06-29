package event

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"mudiralmaham/models"
	"mudiralmaham/utils/database"
	"net/http"
)

func Add(w http.ResponseWriter, r *http.Request) {
	event, err := decodeEvent(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	_, err = database.DB.Collection("event").InsertOne(context.TODO(), event)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var user models.User
	err = database.
		DB.
		Collection("user").
		FindOne(context.TODO(), bson.D{{"username", event.Owner}}).
		Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if user.Username == "" {
		http.Error(w, "event's owner does not exist in database", 400)
		return
	}
	update := bson.M{"$push": bson.M{"events": event}}
	_, err = database.
		DB.
		Collection("user").
		UpdateOne(context.TODO(), bson.D{{"username", event.Owner}}, update)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	_, _ = fmt.Fprint(w, "event added to database")
	w.WriteHeader(200)
}

func Get(w http.ResponseWriter, r *http.Request) {
	me, err := decodeOwner(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var user models.User
	err = database.
		DB.
		Collection("user").FindOne(context.TODO(), bson.M{"username": me.Username}).
		Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	output, err := json.Marshal(user.Events)
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
