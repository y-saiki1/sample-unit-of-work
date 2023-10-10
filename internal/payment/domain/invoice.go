package domain

import (
	"time"
)

type Invoice struct {
	InvoiceId     InvoiceId
	CompanyId     CompanyId
	ClientId      ClientId
	IssueDate     IssueDate
	PaymentAmount PaymentAmount
	Fee           Fee
	Tax           Tax
	InvoiceAmount InvoiceAmount
	DueDate       DueDate

	CurrentStatus InvoiceStatusLog
	StatusLogs    []InvoiceStatusLog
}

func NewInvoice(issueDate, dueDate time.Time, invoiceId, companyId, userId, userName, clientId string, payment int) (Invoice, error) {
	invId, err := NewInvoiceId(invoiceId)
	if err != nil {
		return Invoice{}, err
	}
	comId, err := NewCompanyId(companyId)
	if err != nil {
		return Invoice{}, err
	}
	clId, err := NewClientId(clientId)
	if err != nil {
		return Invoice{}, err
	}

	issDate, err := NewIssueDate(issueDate)
	if err != nil {
		return Invoice{}, err
	}
	dDate, err := NewDueDate(dueDate, issDate)
	if err != nil {
		return Invoice{}, err
	}
	pAmount, err := NewPaymentAmount(payment)
	if err != nil {
		return Invoice{}, err
	}
	statusLog, err := NewInvoiceStatusLog(userId, userName, INVOICE_STATUS_PENDING)
	if err != nil {
		return Invoice{}, err
	}

	return Invoice{
		InvoiceId:     invId,
		ClientId:      clId,
		CompanyId:     comId,
		IssueDate:     issDate,
		DueDate:       dDate,
		PaymentAmount: pAmount,
		CurrentStatus: statusLog,
		StatusLogs:    []InvoiceStatusLog{statusLog},
	}, nil
}

func (i *Invoice) CalculateInvoiceAmount() error {
	i.Fee = NewFee(i.PaymentAmount)
	i.Tax = NewTax(i.Fee)
	invoiceAmount, err := NewInvoiceAmount(i.PaymentAmount, i.Fee, i.Tax)
	if err != nil {
		return err
	}

	i.InvoiceAmount = invoiceAmount
	return nil
}
