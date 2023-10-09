package queryservice

import "errors"

var (
	ErrorFailedToFindUser    error = errors.New("failed to find user. please try again.")
	ErrorFailedToFindInvoces error = errors.New("failed to find invoices. please try again.")
)
