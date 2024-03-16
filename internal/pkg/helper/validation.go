package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidationErrMsg(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "is required"
	case "ascii":
		return "must comply ascii"
	case "email":
		return "must be valid email"
	case "min":
		return fmt.Sprintf("must be atleast %s characters length", e.Param())
	case "eqfield":
		return fmt.Sprintf("must equal to %s field", e.Param())
	case "ulid":
		return "must be ULID value"
	default:
		return "is not valid"
	}
}
