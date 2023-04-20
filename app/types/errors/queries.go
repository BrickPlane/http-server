package errors

import "errors"

var (
	InvalidLogin = errors.New("Login or email allready exist")
	CredError = errors.New("Error in cred")
	InputData = errors.New("Wrong input data")
	Handler = ("Error during data transfer")
)