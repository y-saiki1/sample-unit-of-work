package service

import (
	"upsidr-coding-test/internal/payment/domain"
)

type ClientVerificationService struct {
	logger      Logger
	companyRepo domain.CompanyRepository
	userRepo    domain.UserRepository
}

func NewClientVerificationService(
	logger Logger,
	companyRepo domain.CompanyRepository,
	userRepo domain.UserRepository,
) ClientVerificationService {
	return ClientVerificationService{
		logger:      logger,
		companyRepo: companyRepo,
		userRepo:    userRepo,
	}
}

func (s *ClientVerificationService) VerifyRelationshipWithClient(userId string, clientId string) error {
	clId, err := domain.NewClientId(clientId)
	if err != nil {
		s.logger.Error(err)
		return err
	}
	uId, err := domain.NewUserId(userId)
	if err != nil {
		s.logger.Error(err)
		return err
	}
	user, err := s.userRepo.FindByUserId(uId)
	if err != nil {
		s.logger.Error(err)
		return ErrorClientRelationVerificationFailed
	}
	ok, err := s.companyRepo.IsClientOfCompany(user.CompanyId, clId)
	if err != nil {
		s.logger.Error(err)
		return ErrorClientRelationVerificationFailed
	}
	if !ok {
		return ErrorClientNotRelatedWithCompany
	}
	return nil
}
