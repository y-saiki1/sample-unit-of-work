package handler

import (
	"fmt"
	"net/http"
	"time"
	"upsidr-coding-test/internal/payment/infra"
	"upsidr-coding-test/internal/payment/service"
	"upsidr-coding-test/internal/payment/usecase/queryservice"

	"upsidr-coding-test/internal/payment/usecase"

	"github.com/labstack/echo/v4"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) GetInvoices(ctx echo.Context, params GetInvoicesParams) error {
	u := ctx.Get("user").(User)

	var dueFrom *time.Time
	if params.DueFrom != nil {
		t, err := time.Parse("2006-01-02", *params.DueFrom)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: ErrorParseDueFrom.Error(),
			})
		}
		dueFrom = &t
	}
	var dueTo *time.Time
	if params.DueTo != nil {
		t, err := time.Parse("2006-01-02", *params.DueTo)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: ErrorParseDueTo.Error(),
			})
		}
		dueTo = &t
	}
	var page int
	if params.Page != nil {
		page = *params.Page
	}
	var containsDeleted bool
	if params.ContainsDeleted != nil {
		containsDeleted = *params.ContainsDeleted
	}

	invQuery := infra.NewInvoiceRDBQuery()
	userRepo := infra.NewUserRDB()
	service := queryservice.NewInvoiceQueryService(ctx.Logger(), &invQuery, &userRepo)
	list, err := service.List(queryservice.InvoiceListDTO{
		UserId:          u.UserId,
		DueFrom:         dueFrom,
		DueTo:           dueTo,
		Page:            page,
		ContainsDeleted: containsDeleted,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
	}

	res := make([]InvoiceListResponse, 0, len(list))
	for _, val := range list {
		var dueDate string
		if val.DueDate != nil {
			dueDate = val.DueDate.Format("2006-01-02")
		}
		var issueDate string
		if val.IssueDate != nil {
			issueDate = val.IssueDate.Format("2006-01-02")
		}
		var updatedAt string
		if val.UpdatedAt != nil {
			updatedAt = val.UpdatedAt.Format("2006-01-02")
		}
		var deletedAt string
		if val.DeletedAt != nil {
			deletedAt = val.DeletedAt.Format("2006-01-02")
		}

		logs := make([]StatusLog, 0, len(val.StatusLogs))
		for _, log := range val.StatusLogs {
			var createdAt string
			if log.CreatedAt != nil {
				createdAt = log.CreatedAt.Format("2006-01-02")
			}
			var updatedAt string
			if log.UpdatedAt != nil {
				updatedAt = log.UpdatedAt.Format("2006-01-02")
			}
			var deletedAt string
			if log.DeletedAt != nil {
				deletedAt = log.DeletedAt.Format("2006-01-02")
			}
			logs = append(logs, StatusLog{
				Status:    log.Status,
				UserName:  log.UserName,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
				DeletedAt: deletedAt,
			})
		}

		res = append(res, InvoiceListResponse{
			CompanyId:     val.CompanyId,
			DueDate:       dueDate,
			PaymentAmount: int(val.PaymentAmount),
			Fee:           int(val.Fee),
			FeeRate:       fmt.Sprintf("%d%%", int(val.FeeRate*100)),
			Tax:           int(val.Tax),
			TaxRate:       fmt.Sprintf("%d%%", int(val.TaxRate*100)),
			InvoiceAmount: int(val.InvoiceAmount),
			InvoiceId:     val.InvoiceId,
			IssueDate:     issueDate,
			UpdatedAt:     updatedAt,
			DeletedAt:     deletedAt,
			Client: Client{
				CompanyId: val.Client.CompanyId,
				Name:      val.Client.Name,
			},
			StatusLogs: logs,
		})
	}

	return ctx.JSON(http.StatusOK, GetInvoicesResponse{
		List: res,
	})
}

func (Server) PostInvoice(ctx echo.Context) error {
	u := ctx.Get("user").(User)

	r := PostInvoiceRequest{}
	if err := ctx.Bind(&r); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Message: ErrorRequestBinding.Error(),
		})
	}
	dueAt, err := time.Parse("2006-01-02", r.DueAt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: ErrorParseDueAt.Error(),
		})
	}

	executor := infra.NewExecutorRDB()
	invoiceRepo := infra.NewInvoiceRDB(&executor)
	companyRepo := infra.NewCompanyRDB()
	userRepo := infra.NewUserRDB()
	clientVerificationService := service.NewClientVerificationService(ctx.Logger(), &companyRepo, &userRepo)
	uc := usecase.NewInvoiceUseCase(ctx.Logger(), &invoiceRepo, &userRepo, &executor, &clientVerificationService)

	invoice, err := uc.Create(usecase.InvoiceCreateDTO{
		UserId:        u.UserId,
		PaymentAmount: r.PaymentAmount,
		DueAt:         dueAt,
		ClientId:      r.ClientId,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, PostInvoiceResponse{
		CompanyId:     invoice.CompanyId.Value(),
		InvoiceId:     invoice.InvoiceId.Value(),
		DueDate:       invoice.DueDate.Value().Format("2006-01-02"),
		Fee:           int(invoice.Fee.Value()),
		InvoiceAmount: int(invoice.InvoiceAmount.Value()),
		IssueDate:     invoice.IssueDate.Value().Format("2006-01-02"),
		PaymentAmount: int(invoice.PaymentAmount.Value()),
		Tax:           int(invoice.Tax.Value()),
	})
}
