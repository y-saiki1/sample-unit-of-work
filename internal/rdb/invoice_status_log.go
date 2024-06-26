package rdb

import "time"

type InvoiceStatusLog struct {
	InvoiceId string     `json:"invoiceId" gorm:"primaryKey"`
	UserId    string     `json:"user_id"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	User *User `json:"user" gorm:"foreignKey:UserId;references:UserId"`
}
