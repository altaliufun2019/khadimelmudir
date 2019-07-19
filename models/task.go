package models

type Task struct {
	Name             string `json:"name"`
	Comment          string `json:"comment"`
	CreatedDate      string `json:"created_date"`
	DueDate          string `json:"due_date"`
	NotificationDate string `json:"notification_date"`
	IsOver           bool   `json:"is_over"`
	IsDone           bool   `json:"is_done"`
	Owner            string `json:"owner"`
	Config           string `json:"config"`
	Project          string `json:"project"`
}
