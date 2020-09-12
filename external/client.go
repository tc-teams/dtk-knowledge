package external

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	BaseURL = "http://languages:5000/document"
)

type Client struct {
	*http.Client
}

func (c *Client) Request(r *BotRequest) (BotRequest, error) {
	reqBytes, err := json.Marshal(r)

	if err != nil {
		return BotRequest{},err
	}
	request, err := http.NewRequest(
		http.MethodPost,
		BaseURL,
		bytes.NewBuffer(reqBytes),
	)
	request.Header.Set("Accept", "application/json; charset=utf-8")

	if err != nil {
		return BotRequest{},err
	}
	resp, err := c.Client.Do(request)
	if err != nil {
		return BotRequest{},  err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return BotRequest{}, err
	}
	var response BotRequest
	if err := json.Unmarshal(body, &response); err != nil {
		return BotRequest{},err

	}

	return response, nil
}

//NewClient return a new client instance
func NewClient() *Client {
	return &Client{
		&http.Client{
			Timeout: time.Minute,
		},
	}
}
