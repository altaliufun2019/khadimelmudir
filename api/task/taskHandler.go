package task

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"mudiralmaham/models"
	"mudiralmaham/utils/database"
	"net/http"
)

func Add(w http.ResponseWriter, r *http.Request) {
	task, err := taskDecoder(r)
	if err != nil {
		http.Error(w, "wrong task format", 400)
		return
	}

	_, err = database.DB.Collection("task").InsertOne(context.TODO(), models.Task{
		Name:             task.Name,
		Comment:          task.Comment,
		CreatedDate:      task.CreatedDate,
		DueDate:          task.DueDate,
		NotificationDate: task.NotificationDate,
		IsOver:           task.IsOver,
		IsDone:           task.IsDone,
		Owner:            task.Owner,
		Config:           task.Config,
	})
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	update := bson.M{"$push": bson.M{"tasks": models.Task{
		Name:             task.Name,
		Comment:          task.Comment,
		CreatedDate:      task.CreatedDate,
		DueDate:          task.DueDate,
		NotificationDate: task.NotificationDate,
		IsOver:           task.IsOver,
		IsDone:           task.IsDone,
		Owner:            task.Owner,
		Config:           task.Config,
	}}}
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

func Update(w http.ResponseWriter, r *http.Request) {
	task, err := taskDecoder(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	_, err = database.DB.Collection("task").DeleteOne(context.TODO(), bson.D{{"Name", task.Name}})
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	_, err = database.DB.Collection("task").InsertOne(context.TODO(), models.Task{
		Name:             task.Name,
		Comment:          task.Comment,
		CreatedDate:      task.CreatedDate,
		DueDate:          task.DueDate,
		NotificationDate: task.NotificationDate,
		IsOver:           task.IsOver,
		IsDone:           task.IsDone,
		Owner:            task.Owner,
		Config:           task.Config,
	})
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	_, _ = fmt.Fprint(w, "task updated successfully")
	w.WriteHeader(200)
}
