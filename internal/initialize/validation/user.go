package validation

import (
	"github.com/go-playground/validator/v10"
)

func password(fl validator.FieldLevel) bool {
	value := []byte(fl.Field().String())
	if len(value) < 8 || len(value) > 16 {
		return false
	}

	var numberNum, upperNum, lowerNum int
	for _, v := range value {
		if v >= '0' && v <= '9' {
			numberNum++
		}

		if v >= 'A' && v <= 'Z' {
			upperNum++
		}

		if v >= 'a' && v <= 'z' {
			lowerNum++
		}
	}

	return numberNum > 0 && upperNum > 0 && lowerNum > 0
}
