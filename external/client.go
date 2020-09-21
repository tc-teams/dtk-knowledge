package external

import (
	"bytes"
	"encoding/json"
	"errors"
	wrap "github.com/pkg/errors"
	"net/http"
	"os"
)

type Client struct {
	*http.Client
}

func (c *Client) Request(r *PlnRequest) (*http.Response, error) {

	reqBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBytes).Encode(r)
	if err != nil {
		errInvalidEncode := errors.New("was not possible enconde response body")
		return nil, wrap.Wrap(err, errInvalidEncode.Error())

	}

	request, err := http.NewRequest(
		http.MethodPost,
		os.Getenv("PLN_URL"),
		reqBytes,
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
