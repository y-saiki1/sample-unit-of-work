package rdb

import "time"

type Invoice struct {
	InvoiceId     string     `json:"invoice_id"`
	CompanyId     string     `json:"company_id"`
	ClientId      string     `json:"client_id"`
	PaymentAmount float64    `json:"payment_amount"`
	Fee           float64    `json:"fee"`
	FeeRate       float64    `json:"fee_rate"`
	Tax           float64    `json:"tax"`
	TaxRate       float64    `json:"tax_rate"`
	TotalAmount   float64    `json:"total_amount"`
	DueAt         *time.Time `json:"due_at"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}
