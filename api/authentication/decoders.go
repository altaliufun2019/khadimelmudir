package authentication

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func loginDecoder(r *http.Request) (loginCredentials, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return loginCredentials{}, err
	}
	var credits loginCredentials
	err = json.Unmarshal(body, &credits)
	if err != nil {
		return loginCredentials{}, err
	}
	return credits, nil
}

func signUpDecoder(r *http.Request) (signUpCredits, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return signUpCredits{}, err
	}
	var credits signUpCredits
	err = json.Unmarshal(body, &credits)
	if err != nil {
		return signUpCredits{}, err
	}
	return credits, nil
}
