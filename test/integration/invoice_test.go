package integration

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"
	"upsidr-coding-test/internal/payment/domain"
	"upsidr-coding-test/internal/payment/handler"
	"upsidr-coding-test/internal/payment/service"
	"upsidr-coding-test/internal/rdb"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

var (
	seeder InvoiceTestSeeder
)

type InvoiceTestSeeder struct {
	companies map[string]rdb.Company
	users     map[string]rdb.User
	clients   map[string]rdb.Client
}

func setUpPostInvoice(t *testing.T) {
	company1 := newCompany(t, "company1", "test_company1", "test1", "09000000001", "0000001", "test_address")
	company2 := newCompany(t, "company2", "test_company2", "test2", "09000000002", "0000002", "test_address")
	company3 := newCompany(t, "company3", "test_company3", "test3", "09000000003", "0000003", "test_address")
	user1 := newUser(t, "user1", "company1", "test_user1", "test1@test.com", "password")
	client1 := newClient(t, "company1", "company2")

	seeder = InvoiceTestSeeder{
		companies: map[string]rdb.Company{
			"company1": company1,
			"company2": company2,
			"company3": company3,
		},
		users: map[string]rdb.User{
			"user1": user1,
		},
		clients: map[string]rdb.Client{
			"client1": client1,
		},
	}
	if err := rdb.DB.Create(&[]rdb.Company{company1, company2}).Error; err != nil {
		e.Logger.Fatal(err)
	}
	if err := rdb.DB.Create(&[]rdb.User{user1}).Error; err != nil {
		e.Logger.Fatal(err)
	}
	if err := rdb.DB.Create(&[]rdb.Client{client1}).Error; err != nil {
		e.Logger.Fatal(err)
	}
}

func TestPostInvoice(t *testing.T) {
	setUpPostInvoice(t)
	tests := []struct {
		name     string
		req      handler.PostInvoiceRequest
		want     handler.PostInvoiceResponse
		wantErr  handler.ErrorResponse
		wantCode int
	}{
		{
			name: "正常にリクエストを送った場合",
			req: handler.PostInvoiceRequest{
				ClientId:      "company2",
				DueAt:         time.Now().AddDate(0, 1, 0).Format("2006-01-02"),
				PaymentAmount: 10000,
			},
			want: handler.PostInvoiceResponse{
				CompanyId:     seeder.companies["company1"].CompanyId,
				DueDate:       time.Now().AddDate(0, 1, 0).Format("2006-01-02"),
				IssueDate:     time.Now().Format("2006-01-02"),
				PaymentAmount: 10000,
				Fee:           400,
				Tax:           40,
				InvoiceAmount: 10440,
			},
		},
		{
			name: "支払い期日が過去の日付の場合",
			req: handler.PostInvoiceRequest{
				ClientId:      "company2",
				DueAt:         time.Now().AddDate(0, 0, -1).Format("2006-01-02"),
				PaymentAmount: 10000,
			},
			wantErr: handler.ErrorResponse{
				Message: domain.ErrorDueDateBeforeIssue.Error(),
			},
		},
		{
			name: "支払い金額が負の整数だった場合",
			req: handler.PostInvoiceRequest{
				ClientId:      "company2",
				DueAt:         time.Now().AddDate(0, 1, 0).Format("2006-01-02"),
				PaymentAmount: -10000,
			},
			wantErr: handler.ErrorResponse{
				Message: domain.ErrorInvalidPaymentAmount.Error(),
			},
		},
		{
			name: "請求対象がクライアント関係になかった場合",
			req: handler.PostInvoiceRequest{
				ClientId:      "company3",
				DueAt:         time.Now().AddDate(0, 1, 0).Format("2006-01-02"),
				PaymentAmount: -10000,
			},
			wantErr: handler.ErrorResponse{
				Message: service.ErrorClientNotRelatedWithCompany.Error(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			usr := handler.User{UserId: seeder.users["user1"].UserId}
			ctx, server, recoder := initServer("/api/invoices", tt.req, usr)
			_ = server.PostInvoice(ctx)

			switch recoder.Code {
			case http.StatusOK:
				res := handler.PostInvoiceResponse{}
				err := json.NewDecoder(recoder.Body).Decode(&res)

				assert.Equal(t, nil, err)
				assert.Equal(t, http.StatusOK, recoder.Code)
				opts := cmpopts.IgnoreFields(handler.PostInvoiceResponse{}, "InvoiceId")
				assert.Empty(t, cmp.Diff(tt.want, res, opts), "different:")

			case http.StatusInternalServerError:
				res := handler.ErrorResponse{}
				err := json.NewDecoder(recoder.Body).Decode(&res)
				assert.Equal(t, nil, err)
				assert.Equal(t, http.StatusInternalServerError, recoder.Code)
				assert.Equal(t, tt.wantErr, res)
			default:
				t.Errorf("unexpected status code: %v", recoder.Code)
			}
		})
	}

}

// func TestGetInvoices(t *testing.T) {

// }
