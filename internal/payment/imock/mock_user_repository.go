package imock

import (
	"upsidr-coding-test/internal/payment/domain"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByUserId(userId domain.UserId) (domain.User, error) {
	args := m.Called(userId)
	return args.Get(0).(domain.User), args.Error(1)
}
