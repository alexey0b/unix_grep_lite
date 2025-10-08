package domain

import "errors"

var (
	ErrInvalidContextLength = errors.New("grep: invalid context length argument")
	ErrWrongArgs            = errors.New("grep: wrong arguments")
)
