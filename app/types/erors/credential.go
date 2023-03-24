package erors

import "errors"

var (
	Login = errors.New("Login field is required")
	Pass = errors.New("Password field is required")
	Email = errors.New("Email field is required")
	Id = errors.New("Id field is required")
)