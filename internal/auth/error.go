package auth

import (
	"errors"
	"fmt"
)

var (
	ErrorUserAuthentication error = errors.New("email or password is invalid.")
	ErrorTokenGenerate      error = errors.New("authentication failed. Please try again.")
	ErrorTokenVerification  error = errors.New("authentication failed. Please try again.")
	ErrorPasswordTooShort   error = errors.New("password is too short.")

	ErrorTokenMethod         error = fmt.Errorf("unexpected signing method.")
	ErrorJWTMapClaimsInvalid error = fmt.Errorf("token claims are not of type jwt.MapClaims.")
)
