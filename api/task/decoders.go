package task

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type taskModel struct {
	Name             string `json:"name"`
	Comment          string `json:"comment"`
	CreatedDate      string `json:"created_date"`
	DueDate          string `json:"due_date"`
	NotificationDate string `json:"notification_date"`
	IsOver           bool   `json:"is_over"`
	IsDone           bool   `json:"is_done"`
	Owner            string `json:"owner"`
	Project          string `json:"project"`
	Config           string `json:"config"`
}

func taskDecoder(r *http.Request) (taskModel, error) {
	var task taskModel
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return task, err
	}

	err = json.Unmarshal(body, task)
	if err != nil {
		return task, err
	}
	return task, nil
}
