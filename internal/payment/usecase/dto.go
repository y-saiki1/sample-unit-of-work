package usecase

import "time"

type InvoiceCreateDTO struct {
	UserId        string
	ClientId      string
	PaymentAmount int
	DueAt         time.Time
}
