package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func getErrValidationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s é obrigatório", fe.Field())
	case "url":
		return fmt.Sprintf("%s deve ser uma URL válida", fe.Field())
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
