package models

type Event struct {
	Name               string `json:"name"`
	CreatedDate        string `json:"created_date"`
	DueDate            string `json:"due_date"`
	NotificationConfig string `json:"notification_config"`
	Content            string `json:"content"`
	Description        string `json:"description"`
	Owner              string `json:"owner"`
}
