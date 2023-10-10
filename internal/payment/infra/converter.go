package infra

import (
	"upsidr-coding-test/internal/payment/domain"
	"upsidr-coding-test/internal/rdb"
)

type invoiceConverter struct{}

func (invoiceConverter) ToRDBModel(invoice domain.Invoice) rdb.Invoice {
	due := invoice.DueDate.Value()
	logs := make([]rdb.InvoiceStatusLog, 0, len(invoice.StatusLogs))
	for _, log := range invoice.StatusLogs {
		logs = append(logs, rdb.InvoiceStatusLog{
			InvoiceId: invoice.InvoiceId.Value(),
			UserId:    log.UserId.Value(),
			Status:    log.Status.Value(),
		})
	}
	return rdb.Invoice{
		InvoiceId:     invoice.InvoiceId.Value(),
		CompanyId:     invoice.CompanyId.Value(),
		ClientId:      invoice.ClientId.Value(),
		PaymentAmount: invoice.PaymentAmount.Value(),
		Fee:           invoice.Fee.Value(),
		FeeRate:       domain.FEE_RATE,
		Tax:           invoice.Tax.Value(),
		TaxRate:       domain.TAX_RATE,
		InvoiceAmount: invoice.InvoiceAmount.Value(),
		DueAt:         &due,

		StatusLogs: logs,
	}
}

type userConverter struct{}

func (userConverter) ToEntity(user rdb.User) (domain.User, error) {
	uId, err := domain.NewUserId(user.UserId)
	if err != nil {
		return domain.User{}, err
	}
	compId, err := domain.NewCompanyId(user.CompanyId)
	if err != nil {
		return domain.User{}, err
	}

	uName, err := domain.NewUserName(user.Name)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		UserId:    uId,
		CompanyId: compId,
		Name:      uName,
	}, nil
}
