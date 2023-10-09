package handler

import (
	"net/http"
	"time"
	"upsidr-coding-test/internal/payment/infra"
	"upsidr-coding-test/internal/payment/service"

	"upsidr-coding-test/internal/payment/usecase"

	"github.com/labstack/echo/v4"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) GetInvoices(ctx echo.Context, params GetInvoicesParams) error {
	return nil
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
			Message: ErrorParseTime.Error(),
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
