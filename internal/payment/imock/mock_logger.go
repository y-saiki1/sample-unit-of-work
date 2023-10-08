package imock

import "github.com/stretchr/testify/mock"

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Error(args ...interface{}) {
	m.Called(args...)
}
