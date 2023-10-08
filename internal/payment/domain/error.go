package domain

import "errors"

var (
	ErrorUserIdEmpty    error = errors.New("user ID must not be an empty string.")
	ErrorCompanyIdEmpty error = errors.New("company ID must not be an empty string.")
	ErrorClientIdEmpty  error = errors.New("client ID must not be an empty string.")
	ErrorInvoiceIdEmpty error = errors.New("invoice ID must not be an empty string.")
)
