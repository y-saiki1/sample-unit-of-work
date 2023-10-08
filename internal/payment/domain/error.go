package domain

import "errors"

var (
	ErrorIssueDateInvalid      error = errors.New("issue date must not be a past date and not a future date.")
	ErrorDueDateBeforeIssue    error = errors.New("due date cannot be before the issue date")
	ErrorInvalidPaymentAmount  error = errors.New("payment amount should not be less than 0")
	ErrorNegativeInvoiceAmount error = errors.New("invoice amount or fee or tax value must not be less than 0")
	ErrorUserIdEmpty           error = errors.New("user ID must not be an empty string.")
	ErrorCompanyIdEmpty        error = errors.New("company ID must not be an empty string.")
	ErrorClientIdEmpty         error = errors.New("client ID must not be an empty string.")
	ErrorInvoiceIdEmpty        error = errors.New("invoice ID must not be an empty string.")
)
