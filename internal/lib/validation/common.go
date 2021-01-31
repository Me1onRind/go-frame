package validation

import (
	"github.com/go-playground/validator/v10"
)

func Page(fl validator.FieldLevel) bool {
	page := fl.Field().Int()
	if page < 0 {
		return false
	}
	if page == 0 {
		fl.Field().SetInt(1)
	}
	return true
}

func PageSize(fl validator.FieldLevel) bool {
	pageSize := fl.Field().Int()
	if pageSize < 0 {
		return false
	}
	if pageSize == 0 {
		fl.Field().SetInt(10)
	}
	return true
}
