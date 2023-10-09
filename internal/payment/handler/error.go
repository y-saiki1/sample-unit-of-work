package handler

import "errors"

var (
	ErrorRequestBinding error = errors.New("failed to process request")
	ErrorCreateInvoice  error = errors.New("failed to create invoice")
	ErrorParseDueAt     error = errors.New("failed to parse dueAt")
	ErrorParseDueFrom   error = errors.New("failed to parse dueFrom")
	ErrorParseDueTo     error = errors.New("failed to parse dueTo")
)
