package usecase

import "errors"

var (
	ErrorFailedToIssueInvoice error = errors.New("failed to issue the invoice. please try again.")
)
