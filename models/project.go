package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	Name          string             `json:"name"`
	CreatedDate   primitive.DateTime `json:"created_date"`
	Description   string             `json:"description"`
	Collaborators string             `json:"collaborators"`
	Tasks         []Task             `json:"tasks"`
}
