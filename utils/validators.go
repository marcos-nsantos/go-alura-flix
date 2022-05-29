package utils

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notblank", validators.NotBlank)
	}
}

func getErrValidationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s é obrigatório", fe.Field())
	case "url":
		return fmt.Sprintf("%s deve ser uma URL válida", fe.Field())
	case "notblank":
		return fmt.Sprintf("%s não pode ser vazio", fe.Field())
	case "iscolor":
		return fmt.Sprintf("%s deve estar no formato hexadecimal, RGB, RGBA, HSL ou HSLA", fe.Field())
	}
	return fmt.Sprintf("%s é inválido", fe.Field())
}

func GetErrValidationMessageResponse(err error) string {
	if err == nil {
		return ""
	}
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			return getErrValidationMessage(fe)
		}
	}
	return err.Error()
}
