package utils

import (
	"reflect"
	"strings"

	"github.com/erikyvanov/chat-fh/models"
	"github.com/go-playground/validator/v10"
)

func useJsonTag(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}

	return name
}

func getErrorsFieldsResponse(errors validator.ValidationErrors) []*models.ErrorFieldResponse {
	var totalErrors []*models.ErrorFieldResponse

	for _, err := range errors {
		var element models.ErrorFieldResponse
		element.FailedField = err.Field()
		element.Tag = err.Tag()
		element.Value = err.Param()
		totalErrors = append(totalErrors, &element)
	}

	return totalErrors
}

func ValidateStruct(structToValidate interface{}) []*models.ErrorFieldResponse {
	validate := validator.New()
	validate.RegisterTagNameFunc(useJsonTag)

	err := validate.Struct(structToValidate)
	if err != nil {
		return getErrorsFieldsResponse(err.(validator.ValidationErrors))
	}

	return nil
}

func ValidateLogin(user models.User) []*models.ErrorFieldResponse {
	validate := validator.New()
	validate.RegisterTagNameFunc(useJsonTag)

	err := validate.StructExcept(user, "Name")
	if err != nil {
		return getErrorsFieldsResponse(err.(validator.ValidationErrors))
	}

	return nil
}
