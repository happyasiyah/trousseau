package trousseau

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Token struct {
	Expire string `json:"expire"`
	Token  string `json:"token"`
}

func Authenticate(server string, username string, password string) (Token, error) {
	absoluteURL, err := endpointURL(server, "api/v1/auth")
	if err != nil {
		return Token{}, fmt.Errorf("unable to compute absolute url to remote endpoint: %v\n", err)
	}

	requestData, err := json.Marshal(struct{ Username, Password string }{username, password})
	if err != nil {
		return Token{}, fmt.Errorf("unable to serialize request to server: %v\n", err)
	}

	req, err := http.NewRequest("POST", absoluteURL, strings.NewReader(string(requestData)))
	if err != nil {
		return Token{}, fmt.Errorf("unable to produce request to server: %v\n", err)
	}

	req.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return Token{}, fmt.Errorf("authentication request failed: %v\n", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Token{}, fmt.Errorf("unable to read response data: %v\n", err)
	}

	var token Token
	err = json.Unmarshal(data, &token)
	if err != nil {
		return Token{}, fmt.Errorf("unable to deserialize response data: %v\n", err)
	}

	return token, nil
}

func endpointURL(baseURL, endpoint string) (string, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	resource, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}

	return base.ResolveReference(resource).String(), nil
}
