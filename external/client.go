package external

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	wrap "github.com/pkg/errors"
	"net/http"
	"os"
)

type Client struct {
	*http.Client
}

func (c *Client) Request(r interface{}) (*http.Response, error) {

	reqBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBytes).Encode(r)
	if err != nil {
		errInvalidEncode := errors.New("was not possible enconde response body")
		return nil, wrap.Wrap(err, errInvalidEncode.Error())

	}
	var UrlResult string

	switch r.(type) {
	case ReqDocuments:
		UrlResult = fmt.Sprintf("%s%s",os.Getenv("PLN_URL"),summary)
	default:
		UrlResult = os.Getenv("PLN_URL")
	}
	fmt.Println("ulr request:",UrlResult)

	request, err := http.NewRequest(
		http.MethodPost,
		UrlResult,
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
