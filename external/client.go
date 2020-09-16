package external

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type Client struct {
	*http.Client
}

func (c *Client) Request(r *PlnRequest) (*http.Response, error) {
	reqBytes, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		http.MethodPost,
		os.Getenv("PLN_URL"),
		bytes.NewBuffer(reqBytes),
	)
	request.Header.Set("Accept", "application/json; charset=utf-8")

	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//NewClient return a new client instance
func NewClient() *Client {
	return &Client{
		&http.Client{},
	}
}
