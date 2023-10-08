package infra

import (
	"upsidr-coding-test/internal/payment/domain"

	"gorm.io/gorm"
)

type InvoiceRDB struct {
	executor *ExecutorRDB
}

func NewInvoiceRDB(executor *ExecutorRDB) InvoiceRDB {
	return InvoiceRDB{executor: executor}
}

func (r *InvoiceRDB) Store(invoice domain.Invoice) error {
	model := invoiceConverter{}.ToRDBModel(invoice)
	r.executor.Append(func(tx *gorm.DB) error {
		return tx.Create(&model).Error
	})
	return nil
}
