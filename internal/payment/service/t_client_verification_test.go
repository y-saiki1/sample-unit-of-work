package service

import (
	"errors"
	"testing"
	"upsidr-coding-test/internal/payment/domain"
	"upsidr-coding-test/internal/payment/imock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestVerifyRelationshipWithClient(t *testing.T) {
	// uId, _ := domain.NewUserId("user1")

	type args struct {
		userId   string
		clientId string
	}
	tests := []struct {
		name    string
		args    args
		mock    func(logger *imock.MockLogger, userRepo *imock.MockUserRepository, companyRepo *imock.MockCompanyRepository)
		wantErr error
	}{
		{
			name: "ユーザーとクライアントの関係性が正しい",
			args: args{userId: "user1", clientId: "client1"},
			mock: func(logger *imock.MockLogger, userRepo *imock.MockUserRepository, companyRepo *imock.MockCompanyRepository) {
				comId, _ := domain.NewCompanyId("company1")
				logger.On("Error", nil).Return()
				userRepo.On("FindByUserId", mock.AnythingOfType("domain.UserId")).Return(domain.User{CompanyId: comId}, nil)
				companyRepo.On("IsClientOfCompany", mock.AnythingOfType("domain.CompanyId"), mock.AnythingOfType("domain.ClientId")).Return(true, nil)
			},
			wantErr: nil,
		},
		{
			name: "関係性の確認中にDBエラーが発生",
			args: args{userId: "user1", clientId: "client1"},
			mock: func(logger *imock.MockLogger, userRepo *imock.MockUserRepository, companyRepo *imock.MockCompanyRepository) {
				comId, _ := domain.NewCompanyId("company1")
				logger.On("Error", mock.Anything).Return()
				userRepo.On("FindByUserId", mock.AnythingOfType("domain.UserId")).Return(domain.User{CompanyId: comId}, nil)
				companyRepo.On("IsClientOfCompany", mock.AnythingOfType("domain.CompanyId"), mock.AnythingOfType("domain.ClientId")).Return(false, errors.New("db error"))
			},
			wantErr: ErrorClientRelationVerificationFailed,
		},
		{
			name: "ユーザーがクライアントとの関係性がない",
			args: args{userId: "user1", clientId: "client1"},
			mock: func(logger *imock.MockLogger, userRepo *imock.MockUserRepository, companyRepo *imock.MockCompanyRepository) {
				comId, _ := domain.NewCompanyId("company1")
				logger.On("Error", nil).Return()
				userRepo.On("FindByUserId", mock.AnythingOfType("domain.UserId")).Return(domain.User{CompanyId: comId}, nil)
				companyRepo.On("IsClientOfCompany", mock.AnythingOfType("domain.CompanyId"), mock.AnythingOfType("domain.ClientId")).Return(false, nil)
			},
			wantErr: ErrorClientNotRelatedWithCompany,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			logger := new(imock.MockLogger)
			userRepo := new(imock.MockUserRepository)
			companyRepo := new(imock.MockCompanyRepository)
			tt.mock(logger, userRepo, companyRepo)
			service := NewClientVerificationService(logger, companyRepo, userRepo)
			err := service.VerifyRelationshipWithClient(tt.args.userId, tt.args.clientId)

			assert.Equal(t, tt.wantErr, err)
			userRepo.AssertExpectations(t)
			companyRepo.AssertExpectations(t)
		})
	}
}
