package validation

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Validate validates the incoming request data
func Validate(ctx *gin.Context, data any) []Error {
	validate := ctx.MustGet("validator").(*validator.Validate)
	if err := validate.Struct(data); err != nil {
		errs := err.(validator.ValidationErrors)
		out := make([]Error, len(errs))
		for i, fe := range errs {
			out[i] = Error{fe.Field(), getErrorMsg(fe)}
		}
		return out
	}
	return nil
}

// getErrorMsg returns an error message string depending on the validation error that has occurred
func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "eth_addr":
		return "should be an ETH address"
	case "number":
		return "should be a number"
	}
	return "Unknown error"
}
