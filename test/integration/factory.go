package integration

import (
	"testing"
	"time"
	"upsidr-coding-test/internal/rdb"
)

// type modelFactory struct {
// 	models []any
// }

// func newModelFactory() modelFactory {
// 	return modelFactory{
// 		models: make([]any, 0, 5),
// 	}
// }

// func (f *modelFactory) newCompany(name, representativeName, phoneNumber, postalCode, address string) rdb.Company {
func newCompany(t *testing.T, id, name, representativeName, phoneNumber, postalCode, address string) rdb.Company {
	now := time.Now()
	return rdb.Company{
		CompanyId:          id,
		Name:               name,
		RepresentativeName: representativeName,
		PhoneNumber:        phoneNumber,
		PostalCode:         postalCode,
		Address:            address,
		CreatedAt:          &now,
		UpdatedAt:          &now,
	}
}

// func (f *modelFactory) newUser(userID, companyID, name, email, password string) rdb.User {
func newUser(t *testing.T, id, companyId, name, email, password string) rdb.User {
	now := time.Now()
	return rdb.User{
		UserId:    id,
		CompanyId: companyId,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}

// func (f *modelFactory) newClient(companyID, clientID string) rdb.Client {
func newClient(t *testing.T, companyId, clientId string) rdb.Client {
	return rdb.Client{
		CompanyId: companyId,
		ClientId:  clientId,
	}
}

func newInvoice(t *testing.T, invoiceId, companyId, clientId string, paymentAmount, fee, feeRate, tax, taxRate, invoiceAmount float64, dueAt time.Time) rdb.Invoice {
	now := time.Now()
	return rdb.Invoice{
		InvoiceId:     invoiceId,
		CompanyId:     companyId,
		ClientId:      clientId,
		PaymentAmount: paymentAmount,
		Fee:           fee,
		FeeRate:       feeRate,
		Tax:           tax,
		TaxRate:       taxRate,
		InvoiceAmount: invoiceAmount,
		DueAt:         &dueAt,
		CreatedAt:     &now,
		UpdatedAt:     &now,
		DeletedAt:     nil,
	}
}
