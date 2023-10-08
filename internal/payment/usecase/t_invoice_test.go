package usecase

import (
	"testing"
	"time"
	"upsidr-coding-test/internal/payment/domain"
	"upsidr-coding-test/internal/payment/imock"
	"upsidr-coding-test/internal/payment/service"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

func TestInvoiceUseCase_Create(t *testing.T) {
	now := time.Now()

	type args struct {
		dto InvoiceCreateDTO
	}

	tests := []struct {
		name string
		args args
		mock func(
			logger *imock.MockLogger,
			invoiceRepo *imock.MockInvoiceRepository,
			companyRepo *imock.MockCompanyRepository,
			userRepo *imock.MockUserRepository,
			executor *imock.MockExecutor,
		)
		want    func() domain.Invoice
		wantErr error
	}{
		{
			name: "正常に請求書を作成",
			args: args{
				dto: InvoiceCreateDTO{
					UserId:        "user1",
					ClientId:      "client1",
					PaymentAmount: 10000,
					DueAt:         now.AddDate(0, 0, 1),
				},
			},
			mock: func(
				logger *imock.MockLogger,
				invoiceRepo *imock.MockInvoiceRepository,
				companyRepo *imock.MockCompanyRepository,
				userRepo *imock.MockUserRepository,
				executor *imock.MockExecutor,
			) {
				comId, _ := domain.NewCompanyId("company1")
				logger.On("Error", mock.Anything).Return()
				invoiceRepo.On("Store", mock.AnythingOfType("domain.Invoice")).Return(nil)
				companyRepo.On("IsClientOfCompany", mock.AnythingOfType("domain.CompanyId"), mock.AnythingOfType("domain.ClientId")).Return(true, nil)
				userRepo.On("FindByUserId", mock.AnythingOfType("domain.UserId")).Return(domain.User{CompanyId: comId}, nil)
				executor.On("Exec").Return(nil)
			},
			want: func() domain.Invoice {
				clId, _ := domain.NewClientId("client1")
				paymentAmount, _ := domain.NewPaymentAmount(10000)
				fee := domain.NewFee(paymentAmount)
				tax := domain.NewTax(fee)
				invoiceAmount, _ := domain.NewInvoiceAmount(paymentAmount, fee, tax)
				issDate, _ := domain.NewIssueDate(now)
				dueDate, _ := domain.NewDueDate(now.AddDate(0, 0, 1), issDate)
				return domain.Invoice{
					ClientId:      clId,
					IssueDate:     issDate,
					PaymentAmount: paymentAmount,
					Fee:           fee,
					Tax:           tax,
					InvoiceAmount: invoiceAmount,
					DueDate:       dueDate,
				}
			},
			wantErr: nil,
		},
		{
			name: "支払い金額が0円",
			args: args{
				dto: InvoiceCreateDTO{
					UserId:        "user1",
					ClientId:      "client1",
					PaymentAmount: 0,
					DueAt:         now.AddDate(0, 0, 1),
				},
			},
			mock: func(
				logger *imock.MockLogger,
				invoiceRepo *imock.MockInvoiceRepository,
				companyRepo *imock.MockCompanyRepository,
				userRepo *imock.MockUserRepository,
				executor *imock.MockExecutor,
			) {
				comId, _ := domain.NewCompanyId("company1")
				logger.On("Error", mock.Anything).Return()
				invoiceRepo.On("Store", mock.AnythingOfType("domain.Invoice")).Return(nil)
				companyRepo.On("IsClientOfCompany", mock.AnythingOfType("domain.CompanyId"), mock.AnythingOfType("domain.ClientId")).Return(true, nil)
				userRepo.On("FindByUserId", mock.AnythingOfType("domain.UserId")).Return(domain.User{CompanyId: comId}, nil)
				executor.On("Exec").Return(nil)
			},
			want: func() domain.Invoice {
				return domain.Invoice{}
			},
			wantErr: domain.ErrorInvalidPaymentAmount,
		},
		{
			name: "取引先IDが空",
			args: args{
				dto: InvoiceCreateDTO{
					UserId:        "user1",
					ClientId:      "",
					PaymentAmount: 10000,
					DueAt:         now.AddDate(0, 0, 1),
				},
			},
			mock: func(
				logger *imock.MockLogger,
				invoiceRepo *imock.MockInvoiceRepository,
				companyRepo *imock.MockCompanyRepository,
				userRepo *imock.MockUserRepository,
				executor *imock.MockExecutor,
			) {
				comId, _ := domain.NewCompanyId("company1")
				logger.On("Error", mock.Anything).Return()
				invoiceRepo.On("Store", mock.AnythingOfType("domain.Invoice")).Return(nil)
				companyRepo.On("IsClientOfCompany", mock.AnythingOfType("domain.CompanyId"), mock.AnythingOfType("domain.ClientId")).Return(true, nil)
				userRepo.On("FindByUserId", mock.AnythingOfType("domain.UserId")).Return(domain.User{CompanyId: comId}, nil)
				executor.On("Exec").Return(nil)
			},
			want: func() domain.Invoice {
				return domain.Invoice{}
			},
			wantErr: domain.ErrorClientIdEmpty,
		},
		{
			name: "ユーザーIDが空",
			args: args{
				dto: InvoiceCreateDTO{
					UserId:        "",
					ClientId:      "client1",
					PaymentAmount: 10000,
					DueAt:         now.AddDate(0, 0, 1),
				},
			},
			mock: func(
				logger *imock.MockLogger,
				invoiceRepo *imock.MockInvoiceRepository,
				companyRepo *imock.MockCompanyRepository,
				userRepo *imock.MockUserRepository,
				executor *imock.MockExecutor,
			) {
				comId, _ := domain.NewCompanyId("company1")
				logger.On("Error", mock.Anything).Return()
				invoiceRepo.On("Store", mock.AnythingOfType("domain.Invoice")).Return(nil)
				companyRepo.On("IsClientOfCompany", mock.AnythingOfType("domain.CompanyId"), mock.AnythingOfType("domain.ClientId")).Return(true, nil)
				userRepo.On("FindByUserId", mock.AnythingOfType("domain.UserId")).Return(domain.User{CompanyId: comId}, nil)
				executor.On("Exec").Return(nil)
			},
			want: func() domain.Invoice {
				return domain.Invoice{}
			},
			wantErr: domain.ErrorUserIdEmpty,
		},
		{
			name: "支払い日が過去の日付",
			args: args{
				dto: InvoiceCreateDTO{
					UserId:        "user1",
					ClientId:      "client1",
					PaymentAmount: 10000,
					DueAt:         now.AddDate(0, 0, -1),
				},
			},
			mock: func(
				logger *imock.MockLogger,
				invoiceRepo *imock.MockInvoiceRepository,
				companyRepo *imock.MockCompanyRepository,
				userRepo *imock.MockUserRepository,
				executor *imock.MockExecutor,
			) {
				comId, _ := domain.NewCompanyId("company1")
				logger.On("Error", mock.Anything).Return()
				invoiceRepo.On("Store", mock.AnythingOfType("domain.Invoice")).Return(nil)
				companyRepo.On("IsClientOfCompany", mock.AnythingOfType("domain.CompanyId"), mock.AnythingOfType("domain.ClientId")).Return(true, nil)
				userRepo.On("FindByUserId", mock.AnythingOfType("domain.UserId")).Return(domain.User{CompanyId: comId}, nil)
				executor.On("Exec").Return(nil)
			},
			want: func() domain.Invoice {
				return domain.Invoice{}
			},
			wantErr: domain.ErrorDueDateBeforeIssue,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			logger := new(imock.MockLogger)
			invoiceRepo := new(imock.MockInvoiceRepository)
			companyRepo := new(imock.MockCompanyRepository)
			userRepo := new(imock.MockUserRepository)
			executor := new(imock.MockExecutor)
			tt.mock(logger, invoiceRepo, companyRepo, userRepo, executor)
			clientVerificationService := service.NewClientVerificationService(logger, companyRepo, userRepo)
			uc := NewInvoiceUseCase(logger, invoiceRepo, userRepo, executor, &clientVerificationService)

			got, err := uc.Create(tt.args.dto)
			expected := tt.want()
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, expected.ClientId, got.ClientId)
			assert.Equal(t, expected.PaymentAmount, got.PaymentAmount)
			assert.Equal(t, expected.Fee, got.Fee)
			assert.Equal(t, expected.Tax, got.Tax)
			assert.Equal(t, expected.InvoiceAmount, got.InvoiceAmount)
			assert.Equal(t, expected.DueDate, got.DueDate)
		})
	}
}
