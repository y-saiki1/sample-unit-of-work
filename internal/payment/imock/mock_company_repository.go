package imock

import (
	"upsidr-coding-test/internal/payment/domain"

	"github.com/stretchr/testify/mock"
)

type MockCompanyRepository struct {
	mock.Mock
}

func (m *MockCompanyRepository) IsClientOfCompany(companyId domain.CompanyId, clientId domain.ClientId) (bool, error) {
	args := m.Called(companyId, clientId)
	return args.Bool(0), args.Error(1)
}
