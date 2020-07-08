package api

import (
	"errors"
	"gopkg.in/go-playground/validator.v8"
)

const validation  =  "validator"

//Validate is a struct
type Validation struct {
	Validate *validator.Validate
}

//NewValidate returns a crawler valid of structs
func NewValidate() *Validation {
	config := &validator.Config{
		TagName:validation}

	return &Validation{
		Validate: validator.New(config)}

}

//ValidateStruct V
func (v Validation) ValidateStruct(generic interface{}) (bool, error) {
	err := v.Validate.Struct(generic)
	if err != nil {
		return true, errors.New(err.Error())

	}
	return false, nil

}
