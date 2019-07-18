package project

import (
	"encoding/json"
	"io/ioutil"
	"mudiralmaham/models"
	"net/http"
)

type owner struct {
	Username string `json:"username"`
}

type collaborationRequest struct {
	Username string `json:"username"`
	Project  string `json:"project"`
	Token    string `json:"token"`
}

func decodeProject(r *http.Request) (models.Project, error) {
	var event models.Project
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

func decodeCollaboration(r *http.Request) (collaborationRequest, error) {
	var me collaborationRequest
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
