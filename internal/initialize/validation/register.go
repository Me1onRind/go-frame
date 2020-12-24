package validation

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type registerS struct {
	tag string
	fc  validator.Func
}

var toRegister = []registerS{
	{"passwd", password},
}

func RegisterGinValidation() error {
	engine, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return errors.New("Gin binding.Validator.Engine is not *validator.Validate")
	}
	for _, v := range toRegister {
		if err := engine.RegisterValidation(v.tag, v.fc); err != nil {
			return err
		}
	}
	return nil
}
