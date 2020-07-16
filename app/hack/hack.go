package api

import (
	"encoding/json"
	"net/http"
)

type Hack interface {
	ParseBody(*http.Request) error
}

type Parse struct {}

//ParseBody try to decode the request body into the struct. If there is an error, respond to the client with the error message and a 400 status code.
func (p Parse) ParseBody(i interface{}, r *http.Request) error {

	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		return err
	}
	return nil

}

func NewHack()*Parse{
	return &Parse{}
}