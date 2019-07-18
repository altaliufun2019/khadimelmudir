package task

import (
	"encoding/json"
	"io/ioutil"
	"mudiralmaham/models"
	"net/http"
)

func taskDecoder(r *http.Request) (models.Task, error) {
	var task models.Task
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
