package handler

import "errors"

var (
	ErrorRequestBinding error = errors.New("failed to process request")
	ErrorCreateInvoice  error = errors.New("failed to create invoice")
	ErrorParseTime      error = errors.New("failed to parse dueAt")
)
