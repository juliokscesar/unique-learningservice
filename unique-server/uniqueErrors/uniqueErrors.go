package uniqueErrors

import (
	"errors"
)

var (
	// Failed Checks Errors
	ErrInvalidId    = errors.New("invalid Object ID")
	ErrInvalidEmail = errors.New("invalid email address")

	// Controller Errors
	ErrNotInitialized = errors.New("controller not initialized")

	// User Errors
	ErrEmailRegistered = errors.New("email is already registered")
	ErrInvalidUser     = errors.New("invalid user email")
	ErrInvalidPassword = errors.New("invalid password")
)
