package validation

import (
	"github.com/go-playground/validator/v10"
	"go-frame/internal/constant/user_constant"
	"go-frame/internal/utils/slice"
)

func UserType(fl validator.FieldLevel) bool {
	value := uint8(fl.Field().Uint())
	return slice.InSliceUint8(value, user_constant.AllUserType)
}
