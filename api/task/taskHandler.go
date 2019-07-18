package task

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"mudiralmaham/utils/database"
	"net/http"
)

func Add(w http.ResponseWriter, r *http.Request) {
	task, err := taskDecoder(r)
	if err != nil {
		http.Error(w, "wrong task format", 400)
		return
	}

	_, err = database.DB.Collection("task").InsertOne(context.TODO(), task)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	update := bson.M{"$push": bson.M{"tasks": task}}
	_, err = database.DB.
		Collection("project").
		UpdateOne(context.TODO(), bson.D{{"Name", task.Project}}, update)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	_, _ = fmt.Fprint(w, "task added successfully")
	w.WriteHeader(200)
}
