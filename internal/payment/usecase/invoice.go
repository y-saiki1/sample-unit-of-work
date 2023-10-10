package usecase

import (
	"time"
	"upsidr-coding-test/internal/payment/domain"
	"upsidr-coding-test/internal/payment/service"

	"github.com/google/uuid"
)

type InvoiceUseCase struct {
	logger                    service.Logger
	invoiceRepo               domain.InvoiceRepository
	userRepo                  domain.UserRepository
	executor                  domain.UnitOfWorkExecutor
	clientVerificationService *service.ClientVerificationService
}

func NewInvoiceUseCase(
	logger service.Logger,
	invoiceRepo domain.InvoiceRepository,
	userRepo domain.UserRepository,
	executor domain.UnitOfWorkExecutor,
	clientVerificationService *service.ClientVerificationService,
) InvoiceUseCase {
	return InvoiceUseCase{
		logger:                    logger,
		invoiceRepo:               invoiceRepo,
		userRepo:                  userRepo,
		executor:                  executor,
		clientVerificationService: clientVerificationService,
	}
}

func (u *InvoiceUseCase) Create(dto InvoiceCreateDTO) (domain.Invoice, error) {
	if err := u.clientVerificationService.VerifyRelationshipWithClient(dto.UserId, dto.ClientId); err != nil {
		return domain.Invoice{}, err
	}
	uId, err := domain.NewUserId(dto.UserId)
	if err != nil {
		return domain.Invoice{}, err
	}

	usr, err := u.userRepo.FindByUserId(uId)
	if err != nil {
		u.logger.Error(err)
		return domain.Invoice{}, nil
	}

	issueD := time.Now()
	issueD = time.Date(issueD.Year(), issueD.Month(), issueD.Day(), 0, 0, 0, 0, time.Local)
	invoice, err := domain.NewInvoice(issueD, dto.DueAt, uuid.NewString(), usr.CompanyId.Value(), usr.UserId.Value(), usr.Name.Value(), dto.ClientId, dto.PaymentAmount)
	if err != nil {
		return domain.Invoice{}, err
	}
	if err := invoice.CalculateInvoiceAmount(); err != nil {
		u.logger.Error(err)
		return domain.Invoice{}, err
	}
	if err := u.invoiceRepo.Store(invoice); err != nil {
		u.logger.Error(err)
		return domain.Invoice{}, ErrorFailedToIssueInvoice
	}
	if err := u.executor.Exec(); err != nil {
		u.logger.Error(err)
		return domain.Invoice{}, ErrorFailedToIssueInvoice
	}
	return invoice, nil
}
