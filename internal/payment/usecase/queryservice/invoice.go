package queryservice

import (
	"time"
	"upsidr-coding-test/internal/payment/domain"
	"upsidr-coding-test/internal/payment/domain/querymodel"
	"upsidr-coding-test/internal/payment/service"
)

const (
	InvoiceListLimit = 30
)

type InvoiceQueryService struct {
	logger   service.Logger
	query    querymodel.InvoiceQuery
	userRepo domain.UserRepository
}

func NewInvoiceQueryService(logger service.Logger, query querymodel.InvoiceQuery, userRepo domain.UserRepository) InvoiceQueryService {
	return InvoiceQueryService{logger: logger, query: query, userRepo: userRepo}
}

func (s *InvoiceQueryService) List(dto InvoiceListDTO) ([]querymodel.InvoiceQueryModel, error) {
	uid, err := domain.NewUserId(dto.UserId)
	if err != nil {
		return nil, err
	}
	u, err := s.userRepo.FindByUserId(uid)
	if err != nil {
		s.logger.Error(err)
		return nil, ErrorFailedToFindUser
	}
	if dto.DueFrom != nil {
		t := time.Date(dto.DueFrom.Year(), dto.DueFrom.Month(), dto.DueFrom.Day(), 0, 0, 0, 0, time.Local)
		dto.DueFrom = &t
	}
	if dto.DueTo != nil {
		t := dto.DueTo.AddDate(0, 0, 1)
		dto.DueTo = &t
	}

	filter := querymodel.InvoiceQueryFilter{
		CompanyId:       u.CompanyId.Value(),
		Limit:           InvoiceListLimit,
		ContainsDeleted: dto.ContainsDeleted,
		DueFrom:         dto.DueFrom,
		DueTo:           dto.DueTo,
	}
	if dto.Page > 1 {
		filter.Offset = (dto.Page - 1) * InvoiceListLimit
	}

	list, err := s.query.FindByFilter(filter)
	if err != nil {
		s.logger.Error(err)
		return nil, ErrorFailedToFindInvoces
	}

	return list, nil
}
