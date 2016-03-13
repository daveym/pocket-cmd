package pocket

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// API - Base interface type for Pocket API
type API interface {
	Authenticate(string, interface{}) error
	Authorise(string, string, interface{}) error
	Retrieve(ItemRequest, interface{}) error
}

// Client - Provide access the Pocket API
type Client struct{}

// Authenticate takes the the users consumer key and performs a one time authentication with
// the Pocket API to request access. A Request Token is returned that should be used for all
//  subsequent requests to Pocket.
func (p *Client) Authenticate(consumerKey string, resp interface{}) error {

	request := map[string]string{"consumer_key": consumerKey, "redirect_uri": RedirectURI}
	jsonStr, _ := json.Marshal(request)
	err := postJSON("POST", AuthenticationURL, jsonStr, resp)

	return err
}

// Authorise -  Using the consumerKey and request code, obtain an Access token and Pocket Username
func (p *Client) Authorise(consumerKey string, code string, resp interface{}) error {

	request := map[string]string{"consumer_key": consumerKey, "code": code}
	jsonStr, _ := json.Marshal(request)
	err := postJSON("POST", AuthorisationURL, jsonStr, resp)

	return err
}

// Retrieve -  Pull back items from Pocket
func (p *Client) Retrieve(itemRequest ItemRequest, resp interface{}) error {

	jsonStr, _ := json.Marshal(itemRequest)
	err := postJSON("GET", RetrieveURL, jsonStr, resp)

	return err
}

func postJSON(action string, url string, data []byte, resp interface{}) (err error) {

	req, err := http.NewRequest(action, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF8")
	req.Header.Set("X-Accept", "application/json")

	client := &http.Client{}
	jsonResp, err := client.Do(req)

	return json.NewDecoder(jsonResp.Body).Decode(resp)
}
