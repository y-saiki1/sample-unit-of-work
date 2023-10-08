package imock

import (
	"upsidr-coding-test/internal/payment/domain"

	"github.com/stretchr/testify/mock"
)

type MockInvoiceRepository struct {
	mock.Mock
}

func (m *MockInvoiceRepository) Store(inv domain.Invoice) error {
	args := m.Called(inv)
	return args.Error(0)
}
