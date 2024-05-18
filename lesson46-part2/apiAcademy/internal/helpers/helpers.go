package helpers

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func SetValidatorEngineToUseJSONTags() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

func FillValidationErrorTag(err error, validationErrors map[string]string) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			if _, exists := validationErrors[fe.Field()]; !exists {
				validationErrors[fe.Field()] = ValidationMessageForTag(fe.Tag(), fe.Param())
			}
		}
	}
}

func ValidationMessageForTag(tag string, param any) string {
	switch tag {
	case "required":
		return "Это поле обязательно к заполнению"
	case "email":
		return "Неправильный формат электронной почты"
	case "min":
		return "Длина данного поля не должна быть меньше " + param.(string) + " символов"
	case "max":
		return "Длина данного поля не должна быть больше " + param.(string) + " символов"
	case "unique":
		return "Такое значение уже существует в БД"
	default:
		return ""
	}
}
