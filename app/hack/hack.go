package api

import (
	"encoding/json"
	"net/http"
)

type Parse struct {}

type Hack interface {
	ParseBody(*http.Request) error
	ParseParam(*http.Request) error
}

//ParseBody parse all request api's
func (p Parse) ParseBody(i interface{},r *http.Request) error {

	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		return err
	}
	return nil
}

//ParseParam parse all request api's
func(p Parse) ParseParam(r *http.Request) error {

	return nil
}

func NewHack()*Parse{
	return &Parse{}
}