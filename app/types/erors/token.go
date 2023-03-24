package erors

import "errors"

var (
	NotFound = errors.New("Header is empty")
	Invalid = errors.New("Invalid token")
	Method = errors.New("Method error")
)