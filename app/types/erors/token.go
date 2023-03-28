package erors

import (
	"errors"
)

var (
	NotFound = errors.New("Header is empty")
	NotSame = errors.New("Wrong user")
	Invalid = errors.New("Invalid token")
	Method = errors.New("Method error")
)