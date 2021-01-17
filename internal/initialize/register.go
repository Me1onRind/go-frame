package initialize

import (
	"errors"
	"go-frame/global"
	"go-frame/internal/lib/validation"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type registerS struct {
	tag string
	fc  validator.Func
}

var toRegister = []registerS{
	{"passwd", validation.Password},
}

func RegisterGinValidation() error {
	engine, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return errors.New("Gin binding.Validator.Engine is not *validator.Validate")
	}

	return registerValidation(engine)
}

func RegisterGlobalValidation() error {
	global.Validate = validator.New()
	return registerValidation(global.Validate)
}

func registerValidation(validate *validator.Validate) error {
	for _, v := range toRegister {
		if err := validate.RegisterValidation(v.tag, v.fc); err != nil {
			return err
		}
	}
	return nil
}
