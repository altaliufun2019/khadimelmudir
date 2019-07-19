package project

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/gomail.v2"
	"mudiralmaham/models"
	"mudiralmaham/utils/database"
	"net/http"
	"strings"
)

func AddCollaborator(w http.ResponseWriter, r *http.Request) {
	collaboration, err := decodeCollaboration(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var collaborator models.User
	err = database.DB.Collection("user").
		FindOne(context.TODO(), bson.D{{"username", collaboration.Username}}).
		Decode(&collaborator)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "magaroojoo@gmail.com")
	m.SetHeader("To", collaboration.Username)
	m.SetHeader("Subject", "collaboration request")
	m.SetBody("text/html",
		"<html>"+
			"<head>"+
			"<script src=\"https://unpkg.com/axios/dist/axios.min.js\"></script>"+
			"<script>"+
			"function send() {axios({"+
			"method: post,"+
			"url: 31.184.135.243/project/reject_collaboration/,"+
			"headers:{"+
			"Authorization: Bearer "+collaboration.Token+
			"},"+
			"data: {"+
			"project: "+collaboration.Project+","+
			"username: "+collaboration.Username+","+
			"token: "+collaboration.Token+""+
			"}).then(response => {})}"+
			"</script>"+
			" </head> "+
			"<body>Hi there "+collaborator.Name+"!<br><p>We recently received a request "+
			"from your future partner that you want to be part of the same project, "+
			"if you do not agree with that click reject</p>"+
			"<button type=\"button\" style=\"color: red\" onClick=\"send()\"> reject </body> </html>")
	d := gomail.NewDialer("smtp.gmail.com", 587, "magaroojoo@gmail.com", "majid77??")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func Add(w http.ResponseWriter, r *http.Request) {
	project, err := decodeProject(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	project.Tasks = []models.Task{}
	_, err = database.DB.Collection("project").InsertOne(context.TODO(), project)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	collaborators := strings.Split(project.Collaborators, " __ ")

	for i := 0; i < len(collaborators); i++ {
		var user models.User
		err = database.
			DB.
			Collection("user").
			FindOne(context.TODO(), bson.D{{"username", collaborators[i]}}).
			Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		update := bson.M{"$push": bson.M{"projects": project}}
		_, err = database.
			DB.
			Collection("user").
			UpdateOne(context.TODO(), bson.D{{"username", collaborators[i]}}, update)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	_, _ = fmt.Fprint(w, "project added to database")
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
	output, err := json.Marshal(user.Projects)
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
