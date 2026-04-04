package controllers

import "github.com/GabrielVilarino/gestao-financeira-back.git/exceptions"

func IsValidationError(err error) bool {
	_, ok := err.(*exceptions.ValidationError)
	return ok
}
