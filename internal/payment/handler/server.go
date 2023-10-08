package handler

import (
	"net/http"
	"time"

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
	_ = ctx.Get("user").(User)

	r := PostInvoiceRequest{}
	if err := ctx.Bind(&r); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusBadRequest, ErrorRequestBinding)
	}
	_, err := time.Parse("2006-01-02", r.DueAt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: ErrorParseTime.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, PostInvoiceResponse{
		CompanyId:     "",
		InvoiceId:     "",
		DueDate:       "",
		Fee:           0,
		InvoiceAmount: 0,
		IssueDate:     "",
		PaymentAmount: 0,
		Tax:           0,
	})
}
