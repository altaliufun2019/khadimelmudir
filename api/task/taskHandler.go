package task

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
	task, err := taskDecoder(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var dbTask models.Task
	err = database.DB.Collection("task").
		FindOne(context.TODO(), bson.D{{"name", task.Name}}).Decode(&dbTask)
	if dbTask.Name != "" {
		http.Error(w, "duplicate task", 400)
		return
	}

	_, err = database.DB.Collection("task").InsertOne(context.TODO(), task)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//
	//update := bson.M{"$push": bson.M{"tasks": models.Task{
	//	Name:             task.Name,
	//	Comment:          task.Comment,
	//	CreatedDate:      task.CreatedDate,
	//	DueDate:          task.DueDate,
	//	NotificationDate: task.NotificationDate,
	//	IsOver:           task.IsOver,
	//	IsDone:           task.IsDone,
	//	Owner:            task.Owner,
	//	Config:           task.Config,
	//}}}
	//_, err = database.DB.
	//	Collection("project").
	//	UpdateOne(context.TODO(), bson.D{{"Name", task.Project}}, update)
	//if err != nil {
	//	http.Error(w, err.Error(), 400)
	//	return
	//}
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

	_, err = database.DB.Collection("task").InsertOne(context.TODO(), task)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	_, _ = fmt.Fprint(w, "task updated successfully")
	w.WriteHeader(200)
}

func Get(w http.ResponseWriter, r *http.Request) {
	me, err := decodeProjectOwner(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var tasks []models.Task
	cursor, err := database.
		DB.
		Collection("task").
		Find(context.TODO(), bson.M{"project": me.Project})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	for cursor.Next(context.TODO()) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	output, err := json.Marshal(tasks)
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
