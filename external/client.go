package external

import (
	"net/http"
	"time"
)

const (
	BaseURL = "http://languages:5000/document"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func(c *Client) Request(r *http.Request, v interface{}) error{
    c.RContentType(r)
	request, err := c.HTTPClient.Do(r)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if request.StatusCode != http.StatusOK {
		return err
	}
	return nil
}
func (c * Client) RContentType(r *http.Request) {
	r.Header.Set("Accept", "application/json; charset=utf-8")
}

//NewClient return a new client instance
func NewClient() *Client {
	return &Client{
		BaseURL: BaseURL,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}