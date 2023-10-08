package imock

import "github.com/stretchr/testify/mock"

type MockExecutor struct {
	mock.Mock
}

func (m *MockExecutor) Exec() error {
	args := m.Called()
	return args.Error(0)
}
