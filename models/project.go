package models

type Project struct {
	Name          string `json:"name"`
	CreatedDate   string `json:"created_date"`
	Description   string `json:"description"`
	Collaborators string `json:"collaborators"`
	Tasks         string `json:"tasks"`
}
