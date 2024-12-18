package seliazers

import "github.com/go-playground/validator/v10"

var walletName validator.Func = func(fl validator.FieldLevel) bool {
	// check apakah name null atau berisi
	if fl.Field().String() == "" {
		return false
	}
	return true
}