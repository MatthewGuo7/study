package model

import (
	"github.com/go-playground/validator/v10"
)

func TopicUrl(fl validator.FieldLevel) bool {

	url := fl.Field().String()

	return len(url) > 4
}
