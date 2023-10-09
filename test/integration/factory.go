package integration

import (
	"testing"
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
	return rdb.Company{
		CompanyId:          id,
		Name:               name,
		RepresentativeName: representativeName,
		PhoneNumber:        phoneNumber,
		PostalCode:         postalCode,
		Address:            address,
	}
}

// func (f *modelFactory) newUser(userID, companyID, name, email, password string) rdb.User {
func newUser(t *testing.T, id, companyId, name, email, password string) rdb.User {
	return rdb.User{
		UserId:    id,
		CompanyId: companyId,
		Name:      name,
		Email:     email,
		Password:  password,
	}
}

// func (f *modelFactory) newClient(companyID, clientID string) rdb.Client {
func newClient(t *testing.T, companyId, clientId string) rdb.Client {
	return rdb.Client{
		CompanyId: companyId,
		ClientId:  clientId,
	}
}
