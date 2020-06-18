package valid

import (
	"errors"
	"gopkg.in/go-playground/validator.v8"
)

//Validate is a struct
type Validation struct {
	Valid *validator.Validate
}

//NewValidate returns a news valid of structs
func NewValidate(Name string) *Validation{
	config := &validator.Config{
		TagName: Name,}

	return &Validation{
		Valid: validator.New(config)}

}


func (v Validation) ValidateStruct(generic interface{}) (bool, error) {
	err := v.Valid.Struct(generic)
	if err != nil {
		return true, errors.New(err.Error())

	}
	return false, nil

}
