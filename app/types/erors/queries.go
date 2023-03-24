package erors

import "errors"

var (
	InvalidLogin = errors.New("Login or email allready exist")
	CredError = errors.New("Error in cred")
)