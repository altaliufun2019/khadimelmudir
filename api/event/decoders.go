package event

import (
	"encoding/json"
	"io/ioutil"
	"mudiralmaham/models"
	"net/http"
)

type owner struct {
	Username string `json:"username"`
}

func decodeEvent(r *http.Request) (models.Event, error) {
	var event models.Event
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return event, err
	}

	err = json.Unmarshal(body, &event)
	if err != nil {
		return event, err
	}
	return event, nil
}

func decodeOwner(r *http.Request) (owner, error) {
	var me owner
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return me, err
	}

	err = json.Unmarshal(body, &me)
	if err != nil {
		return me, err
	}
	return me, nil
}
