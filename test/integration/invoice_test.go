package integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
	"upsidr-coding-test/internal/auth"
	"upsidr-coding-test/internal/payment/domain"
	"upsidr-coding-test/internal/payment/handler"
	"upsidr-coding-test/internal/payment/service"
	"upsidr-coding-test/internal/rdb"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type InvoiceTestSeeder struct {
	companies map[string]rdb.Company
	users     map[string]rdb.User
	clients   map[string]rdb.Client
	invoices  map[string]rdb.Invoice
}

func hashed(rawPass string) string {
	h, err := hashPassword(rawPass)
	if err != nil {
		e.Logger.Fatal(err)
	}
	return h
}

func setUpPostInvoice(t *testing.T) InvoiceTestSeeder {
	company1 := newCompany(t, uuid.NewString(), "test_company1", "test1", "09000000001", "0000001", "test_address")
	company2 := newCompany(t, uuid.NewString(), "test_company2", "test2", "09000000002", "0000002", "test_address")
	company3 := newCompany(t, uuid.NewString(), "test_company3", "test3", "09000000003", "0000003", "test_address")
	user1 := newUser(t, uuid.NewString(), company1.CompanyId, "test_user1", fmt.Sprintf("%s@test.com", uuid.NewString()), hashed("password"))
	client1 := newClient(t, company1.CompanyId, company2.CompanyId)

	seeder := InvoiceTestSeeder{
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
	return seeder
}

func setUpGetInvoices(t *testing.T) InvoiceTestSeeder {
	companies := []rdb.Company{
		newCompany(t, uuid.NewString(), "test_company1", "test1", "09000000001", "0000001", "test_address"),
		newCompany(t, uuid.NewString(), "test_company2", "test2", "09000000002", "0000002", "test_address"),
		newCompany(t, uuid.NewString(), "test_company3", "test3", "09000000003", "0000003", "test_address"),
		newCompany(t, uuid.NewString(), "test_company4", "test4", "09000000004", "0000004", "test_address"),
		newCompany(t, uuid.NewString(), "test_company5", "test5", "09000000005", "0000005", "test_address"),
		newCompany(t, uuid.NewString(), "test_company6", "test6", "09000000006", "0000006", "test_address"),
	}
	users := []rdb.User{
		newUser(t, uuid.NewString(), companies[0].CompanyId, hashed("test_user1"), fmt.Sprintf("%s@test.com", uuid.NewString()), "password"),
		newUser(t, uuid.NewString(), companies[1].CompanyId, hashed("test_user2"), fmt.Sprintf("%s@test.com", uuid.NewString()), "password"),
		newUser(t, uuid.NewString(), companies[2].CompanyId, hashed("test_user3"), fmt.Sprintf("%s@test.com", uuid.NewString()), "password"),
		newUser(t, uuid.NewString(), companies[3].CompanyId, hashed("test_user4"), fmt.Sprintf("%s@test.com", uuid.NewString()), "password"),
	}
	clients := []rdb.Client{
		newClient(t, companies[0].CompanyId, companies[1].CompanyId),
		newClient(t, companies[0].CompanyId, companies[2].CompanyId),
		newClient(t, companies[0].CompanyId, companies[3].CompanyId),
		newClient(t, companies[0].CompanyId, companies[4].CompanyId),
		newClient(t, companies[0].CompanyId, companies[5].CompanyId),
	}

	paymentAmounts := []float64{10000, 25000, 50000, 70000, 110000}
	dueAts := []time.Time{time.Now().AddDate(0, 0, 1), time.Now().AddDate(0, 0, 3), time.Now().AddDate(0, 0, 13), time.Now().AddDate(0, 0, 40), time.Now().AddDate(0, 0, 100)}
	invoices := make([]rdb.Invoice, 0, len(paymentAmounts))
	invoiceMP := make(map[string]rdb.Invoice, len(paymentAmounts))
	for i, val := range paymentAmounts {
		idx := i + 1
		fee := (val * domain.FEE_RATE)
		tax := fee * domain.TAX_RATE
		invAmount := val + fee + tax
		dueAt := dueAts[i]
		inv := newInvoice(t, uuid.NewString(), companies[0].CompanyId, users[0].UserId, companies[idx].CompanyId, val, fee, domain.FEE_RATE, tax, domain.TAX_RATE, invAmount, dueAt)
		inv.Client = &companies[idx]
		invoices = append(invoices, inv)
		invoiceMP[fmt.Sprintf("invoice%d", idx)] = inv
	}

	seeder := InvoiceTestSeeder{
		companies: map[string]rdb.Company{
			"company1": companies[0],
			"company2": companies[1],
			"company3": companies[2],
			"company4": companies[3],
			"company5": companies[4],
			"company6": companies[5],
		},
		users: map[string]rdb.User{
			"user1": users[0],
			"user2": users[1],
			"user3": users[2],
			"user4": users[3],
		},
		clients: map[string]rdb.Client{
			"client1": clients[0],
			"client2": clients[1],
			"client3": clients[2],
			"client4": clients[3],
			"client5": clients[4],
		},
		invoices: invoiceMP,
	}
	if err := rdb.DB.Create(&companies).Error; err != nil {
		e.Logger.Fatal(err)
	}
	if err := rdb.DB.Create(&users).Error; err != nil {
		e.Logger.Fatal(err)
	}
	if err := rdb.DB.Create(&clients).Error; err != nil {
		e.Logger.Fatal(err)
	}
	if err := rdb.DB.Debug().Create(&invoices).Error; err != nil {
		e.Logger.Fatal(err)
	}
	return seeder
}

func TestPostInvoice(t *testing.T) {
	seed := setUpPostInvoice(t)
	tk, _ := getToken(seed.users["user1"].UserId)
	tests := []struct {
		name     string
		token    string
		req      handler.PostInvoiceRequest
		want     handler.PostInvoiceResponse
		wantErr  handler.ErrorResponse
		wantCode int
	}{
		{
			name:  "正常にリクエストを送った場合",
			token: tk,
			req: handler.PostInvoiceRequest{
				ClientId:      seed.companies["company2"].CompanyId,
				DueAt:         time.Now().AddDate(0, 1, 0).Format("2006-01-02"),
				PaymentAmount: 10000,
			},
			want: handler.PostInvoiceResponse{
				CompanyId:     seed.companies["company1"].CompanyId,
				DueDate:       time.Now().AddDate(0, 1, 0).Format("2006-01-02"),
				IssueDate:     time.Now().Format("2006-01-02"),
				PaymentAmount: 10000,
				Fee:           400,
				Tax:           40,
				InvoiceAmount: 10440,
			},
		},
		{
			name:  "支払い期日が過去の日付の場合",
			token: tk,
			req: handler.PostInvoiceRequest{
				ClientId:      seed.companies["company2"].CompanyId,
				DueAt:         time.Now().AddDate(0, 0, -1).Format("2006-01-02"),
				PaymentAmount: 10000,
			},
			wantErr: handler.ErrorResponse{
				Message: domain.ErrorDueDateBeforeIssue.Error(),
			},
		},
		{
			name:  "支払い金額が負の整数だった場合",
			token: tk,
			req: handler.PostInvoiceRequest{
				ClientId:      seed.companies["company2"].CompanyId,
				DueAt:         time.Now().AddDate(0, 1, 0).Format("2006-01-02"),
				PaymentAmount: -10000,
			},
			wantErr: handler.ErrorResponse{
				Message: domain.ErrorInvalidPaymentAmount.Error(),
			},
		},
		{
			name:  "請求対象がクライアント関係になかった場合",
			token: tk,
			req: handler.PostInvoiceRequest{
				ClientId:      seed.companies["company3"].CompanyId,
				DueAt:         time.Now().AddDate(0, 1, 0).Format("2006-01-02"),
				PaymentAmount: -10000,
			},
			wantErr: handler.ErrorResponse{
				Message: service.ErrorClientNotRelatedWithCompany.Error(),
			},
		},
		{
			name:  "トークンが無効な場合",
			token: "abcdefg",
			req: handler.PostInvoiceRequest{
				ClientId:      seed.companies["company2"].CompanyId,
				DueAt:         time.Now().AddDate(0, 1, 0).Format("2006-01-02"),
				PaymentAmount: 10000,
			},
			wantErr: handler.ErrorResponse{
				Message: auth.ErrorTokenVerification.Error(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			body, statusCode, err := postRequest("/api/invoices", "POST", tt.req, tt.token)
			assert.Equal(t, nil, err)

			switch statusCode {
			case http.StatusOK:
				res := handler.PostInvoiceResponse{}
				err := json.Unmarshal(body, &res)
				assert.Equal(t, nil, err)
				assert.Equal(t, http.StatusOK, statusCode)
				opts := cmpopts.IgnoreFields(handler.PostInvoiceResponse{}, "InvoiceId")
				assert.Empty(t, cmp.Diff(tt.want, res, opts), "different:")

			case http.StatusInternalServerError:
				res := handler.ErrorResponse{}
				err := json.Unmarshal(body, &res)
				assert.Equal(t, nil, err)
				assert.Equal(t, http.StatusInternalServerError, statusCode)
				assert.Equal(t, tt.wantErr, res)

			case http.StatusUnauthorized:
				res := handler.ErrorResponse{}
				err := json.Unmarshal(body, &res)
				assert.Equal(t, nil, err)

				assert.Equal(t, http.StatusUnauthorized, statusCode)
				assert.Equal(t, tt.wantErr, res)
			default:
				t.Errorf("unexpected status code: %v", statusCode)
			}
		})
	}

}

func TestGetInvoices(t *testing.T) {
	seed := setUpGetInvoices(t)
	tk, _ := getToken(seed.users["user1"].UserId)
	tests := []struct {
		name    string
		token   string
		req     func() map[string]string
		want    func() handler.GetInvoicesResponse
		wantErr handler.ErrorResponse
	}{
		{
			name:  "何もパラメータを指定せず送った場合",
			token: tk,
			req:   func() map[string]string { return nil },
			want: func() handler.GetInvoicesResponse {
				responses := make([]handler.InvoiceListResponse, 0, len(seed.invoices))
				for _, invoice := range seed.invoices {
					responses = append(responses, handler.InvoiceListResponse{
						InvoiceId:     invoice.InvoiceId,
						CompanyId:     invoice.CompanyId,
						DueDate:       invoice.DueAt.Format("2006-01-02"),
						IssueDate:     invoice.CreatedAt.Format("2006-01-02"),
						PaymentAmount: int(invoice.PaymentAmount),
						Fee:           int(invoice.Fee),
						FeeRate:       fmt.Sprintf("%d%%", int(invoice.FeeRate*100)),
						Tax:           int(invoice.Tax),
						TaxRate:       fmt.Sprintf("%d%%", int(invoice.TaxRate*100)),
						InvoiceAmount: int(invoice.InvoiceAmount),
						UpdatedAt:     invoice.UpdatedAt.Format("2006-01-02"),
						DeletedAt:     "",

						Client: handler.Client{
							CompanyId: invoice.Client.CompanyId,
							Name:      invoice.Client.Name,
						},
						StatusLogs: []handler.StatusLog{{
							Status:    string(domain.INVOICE_STATUS_PENDING),
							UserName:  seed.users["user1"].Name,
							CreatedAt: invoice.CreatedAt.Format("2006-01-02"),
							UpdatedAt: invoice.UpdatedAt.Format("2006-01-02"),
						}},
					})
				}
				return handler.GetInvoicesResponse{List: responses}
			},
			wantErr: handler.ErrorResponse{},
		},
		{
			name:  "1ヶ月先からの請求データを確認した場合",
			token: tk,
			req: func() map[string]string {
				dueFrom := time.Now().AddDate(0, 1, 0).Format("2006-01-02")
				return map[string]string{
					"dueFrom": dueFrom,
				}
			},
			want: func() handler.GetInvoicesResponse {
				responses := make([]handler.InvoiceListResponse, 0, 3)
				for _, invoice := range []rdb.Invoice{seed.invoices["invoice4"], seed.invoices["invoice5"]} {
					responses = append(responses, handler.InvoiceListResponse{
						InvoiceId:     invoice.InvoiceId,
						CompanyId:     invoice.CompanyId,
						DueDate:       invoice.DueAt.Format("2006-01-02"),
						IssueDate:     invoice.CreatedAt.Format("2006-01-02"),
						PaymentAmount: int(invoice.PaymentAmount),
						Fee:           int(invoice.Fee),
						FeeRate:       fmt.Sprintf("%d%%", int(invoice.FeeRate*100)),
						Tax:           int(invoice.Tax),
						TaxRate:       fmt.Sprintf("%d%%", int(invoice.TaxRate*100)),
						InvoiceAmount: int(invoice.InvoiceAmount),
						UpdatedAt:     invoice.UpdatedAt.Format("2006-01-02"),
						DeletedAt:     "",

						Client: handler.Client{
							CompanyId: invoice.Client.CompanyId,
							Name:      invoice.Client.Name,
						},
						StatusLogs: []handler.StatusLog{{
							Status:    string(domain.INVOICE_STATUS_PENDING),
							UserName:  seed.users["user1"].Name,
							CreatedAt: invoice.CreatedAt.Format("2006-01-02"),
							UpdatedAt: invoice.UpdatedAt.Format("2006-01-02"),
						}},
					})
				}

				return handler.GetInvoicesResponse{List: responses}
			},
			wantErr: handler.ErrorResponse{},
		},
		{
			name:  "3ヶ月先からの請求データを確認した場合",
			token: tk,
			req: func() map[string]string {
				dueFrom := time.Now().AddDate(0, 3, 0).Format("2006-01-02")
				return map[string]string{
					"dueFrom": dueFrom,
				}
			},
			want: func() handler.GetInvoicesResponse {
				responses := make([]handler.InvoiceListResponse, 0, 3)
				for _, invoice := range []rdb.Invoice{seed.invoices["invoice5"]} {
					responses = append(responses, handler.InvoiceListResponse{
						InvoiceId:     invoice.InvoiceId,
						CompanyId:     invoice.CompanyId,
						DueDate:       invoice.DueAt.Format("2006-01-02"),
						IssueDate:     invoice.CreatedAt.Format("2006-01-02"),
						PaymentAmount: int(invoice.PaymentAmount),
						Fee:           int(invoice.Fee),
						FeeRate:       fmt.Sprintf("%d%%", int(invoice.FeeRate*100)),
						Tax:           int(invoice.Tax),
						TaxRate:       fmt.Sprintf("%d%%", int(invoice.TaxRate*100)),
						InvoiceAmount: int(invoice.InvoiceAmount),
						UpdatedAt:     invoice.UpdatedAt.Format("2006-01-02"),
						DeletedAt:     "",

						Client: handler.Client{
							CompanyId: invoice.Client.CompanyId,
							Name:      invoice.Client.Name,
						},
						StatusLogs: []handler.StatusLog{{
							Status:    string(domain.INVOICE_STATUS_PENDING),
							UserName:  seed.users["user1"].Name,
							CreatedAt: invoice.CreatedAt.Format("2006-01-02"),
							UpdatedAt: invoice.UpdatedAt.Format("2006-01-02"),
						}},
					})
				}

				return handler.GetInvoicesResponse{List: responses}
			},
			wantErr: handler.ErrorResponse{},
		},
		{
			name:  "現在から3日以内の請求データを確認した場合",
			token: tk,
			req: func() map[string]string {
				dueFrom := time.Now().Format("2006-01-02")
				dueTo := time.Now().AddDate(0, 0, 3).Format("2006-01-02")
				return map[string]string{
					"dueFrom": dueFrom,
					"dueTo":   dueTo,
				}
			},
			want: func() handler.GetInvoicesResponse {
				responses := make([]handler.InvoiceListResponse, 0, 3)
				for _, invoice := range []rdb.Invoice{seed.invoices["invoice1"], seed.invoices["invoice2"]} {
					responses = append(responses, handler.InvoiceListResponse{
						InvoiceId:     invoice.InvoiceId,
						CompanyId:     invoice.CompanyId,
						DueDate:       invoice.DueAt.Format("2006-01-02"),
						IssueDate:     invoice.CreatedAt.Format("2006-01-02"),
						PaymentAmount: int(invoice.PaymentAmount),
						Fee:           int(invoice.Fee),
						FeeRate:       fmt.Sprintf("%d%%", int(invoice.FeeRate*100)),
						Tax:           int(invoice.Tax),
						TaxRate:       fmt.Sprintf("%d%%", int(invoice.TaxRate*100)),
						InvoiceAmount: int(invoice.InvoiceAmount),
						UpdatedAt:     invoice.UpdatedAt.Format("2006-01-02"),
						DeletedAt:     "",

						Client: handler.Client{
							CompanyId: invoice.Client.CompanyId,
							Name:      invoice.Client.Name,
						},
						StatusLogs: []handler.StatusLog{{
							Status:    string(domain.INVOICE_STATUS_PENDING),
							UserName:  seed.users["user1"].Name,
							CreatedAt: invoice.CreatedAt.Format("2006-01-02"),
							UpdatedAt: invoice.UpdatedAt.Format("2006-01-02"),
						}},
					})
				}

				return handler.GetInvoicesResponse{List: responses}
			},
			wantErr: handler.ErrorResponse{},
		},
		{
			name:  "トークンが無効な場合",
			token: "abcdefg",
			req: func() map[string]string {
				dueFrom := time.Now().Format("2006-01-02")
				dueTo := time.Now().AddDate(0, 0, 3).Format("2006-01-02")
				return map[string]string{
					"dueFrom": dueFrom,
					"dueTo":   dueTo,
				}
			},
			wantErr: handler.ErrorResponse{
				Message: auth.ErrorTokenVerification.Error(),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := tt.req()
			body, statusCode, err := getRequest("/api/invoices", "GET", req, tt.token)
			assert.Equal(t, err, nil)

			switch statusCode {
			case http.StatusOK:
				res := handler.GetInvoicesResponse{}
				want := tt.want()
				err := json.Unmarshal(body, &res)
				assert.Equal(t, nil, err)

				assert.Equal(t, nil, err)
				assert.Equal(t, http.StatusOK, statusCode)
				assert.ElementsMatch(t, want.List, res.List)

			case http.StatusInternalServerError:
				res := handler.ErrorResponse{}
				err := json.Unmarshal(body, &res)
				assert.Equal(t, nil, err)

				assert.Equal(t, nil, err)
				assert.Equal(t, http.StatusInternalServerError, statusCode)
				assert.Equal(t, tt.wantErr, res)

			case http.StatusUnauthorized:
				res := handler.ErrorResponse{}
				err := json.Unmarshal(body, &res)
				assert.Equal(t, nil, err)

				assert.Equal(t, http.StatusUnauthorized, statusCode)
				assert.Equal(t, tt.wantErr, res)
			default:
				t.Errorf("unexpected status code: %v", statusCode)
			}
		})
	}
}
