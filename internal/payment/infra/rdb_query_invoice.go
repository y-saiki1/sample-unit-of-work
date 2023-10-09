package infra

import (
	"upsidr-coding-test/internal/payment/domain/querymodel"
	"upsidr-coding-test/internal/rdb"

	"gorm.io/gorm"
)

type InvoiceRDBQuery struct{}

func NewInvoiceRDBQuery() InvoiceRDBQuery {
	return InvoiceRDBQuery{}
}

func (q *InvoiceRDBQuery) FindByFilter(filter querymodel.InvoiceQueryFilter) ([]querymodel.InvoiceQueryModel, error) {
	query := rdb.DB.Table("invoices")
	if filter.DueFrom != nil {
		query.Where("invoices.due_at >= ?", filter.DueFrom)
	}
	if filter.DueTo != nil {
		query.Where("invoices.due_at < ?", filter.DueTo)
	}
	if !filter.ContainsDeleted {
		query.Where("invoices.deleted_at is null")
	}

	invoices := []rdb.Invoice{}
	if err := query.
		Where("company_id = ?", filter.CompanyId).
		Preload("StatusLogs", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC")
		}).
		Preload("StatusLogs.User").
		Preload("Client").
		Offset(filter.Offset).
		Limit(filter.Limit).
		Find(&invoices).Error; err != nil {
		return nil, err
	}

	list := make([]querymodel.InvoiceQueryModel, 0, len(invoices))
	for _, invoice := range invoices {
		logs := make([]querymodel.InvoiceStatusLogQueryModel, 0, len(invoice.StatusLogs))
		for _, log := range invoice.StatusLogs {
			logs = append(logs, querymodel.InvoiceStatusLogQueryModel{
				UserName:  log.User.Name,
				Status:    log.Status,
				CreatedAt: log.CreatedAt,
				UpdatedAt: log.UpdatedAt,
				DeletedAt: log.DeletedAt,
			})
		}

		list = append(list, querymodel.InvoiceQueryModel{
			InvoiceId:     invoice.InvoiceId,
			CompanyId:     invoice.CompanyId,
			PaymentAmount: invoice.PaymentAmount,
			Fee:           invoice.Fee,
			FeeRate:       invoice.FeeRate,
			Tax:           invoice.Tax,
			TaxRate:       invoice.TaxRate,
			InvoiceAmount: invoice.InvoiceAmount,
			DueDate:       invoice.DueAt,
			IssueDate:     invoice.CreatedAt,
			UpdatedAt:     invoice.UpdatedAt,
			DeletedAt:     invoice.DeletedAt,

			Client: querymodel.InvoiceClientQueryModel{
				CompanyId: invoice.Client.CompanyId,
				Name:      invoice.Client.Name,
			},
			StatusLogs: logs,
		})
	}

	return list, nil
}
