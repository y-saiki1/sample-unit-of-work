package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewInvoice(t *testing.T) {
	issueDate := time.Now()
	dueDate := time.Now().AddDate(0, 0, 1)
	paymentAmount := 10000.00

	type args struct {
		issueDate time.Time
		dueDate   time.Time
		invoiceId string
		companyId string
		clientId  string
		payment   int
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
		want    Invoice
	}{
		{
			name: "正常なデータ",
			args: args{
				issueDate: issueDate,
				dueDate:   dueDate,
				invoiceId: "invoice1",
				companyId: "company1",
				clientId:  "client1",
				payment:   int(paymentAmount),
			},
			want: Invoice{
				InvoiceId:     InvoiceId{"invoice1"},
				CompanyId:     CompanyId{"company1"},
				ClientId:      ClientId{"client1"},
				IssueDate:     IssueDate{issueDate},
				PaymentAmount: PaymentAmount{paymentAmount},
				DueDate:       DueDate{dueDate},
			},
			wantErr: nil,
		},
		{
			name: "支払額が負の値",
			args: args{
				issueDate: issueDate,
				dueDate:   dueDate,
				invoiceId: "invoice1",
				companyId: "company1",
				clientId:  "client1",
				payment:   -int(paymentAmount),
			},
			want:    Invoice{},
			wantErr: ErrorInvalidPaymentAmount,
		},
		{
			name: "請求書IDが空",
			args: args{
				issueDate: issueDate,
				dueDate:   dueDate,
				invoiceId: "",
				companyId: "company1",
				clientId:  "client1",
				payment:   int(paymentAmount),
			},
			want:    Invoice{},
			wantErr: ErrorInvoiceIdEmpty,
		},
		{
			name: "会社IDが空",
			args: args{
				issueDate: issueDate,
				dueDate:   dueDate,
				invoiceId: "invoice1",
				companyId: "",
				clientId:  "client1",
				payment:   int(paymentAmount),
			},
			want:    Invoice{},
			wantErr: ErrorCompanyIdEmpty,
		},
		{
			name: "取引先IDが空",
			args: args{
				issueDate: issueDate,
				dueDate:   dueDate,
				invoiceId: "invoice1",
				companyId: "company1",
				clientId:  "",
				payment:   int(paymentAmount),
			},
			want:    Invoice{},
			wantErr: ErrorClientIdEmpty,
		},
		{
			name: "発行日が支払日よりも前",
			args: args{
				issueDate: issueDate,
				dueDate:   dueDate,
				invoiceId: "invoice1",
				companyId: "company1",
				clientId:  "client1",
				payment:   int(paymentAmount),
			},
			want: Invoice{
				InvoiceId:     InvoiceId{"invoice1"},
				CompanyId:     CompanyId{"company1"},
				ClientId:      ClientId{"client1"},
				IssueDate:     IssueDate{issueDate},
				PaymentAmount: PaymentAmount{paymentAmount},
				// Fee:           Fee{fee},
				// Tax:           Tax{tax},
				// InvoiceAmount: InvoiceAmount{invoiceAmount},
				DueDate: DueDate{dueDate},
			},
			wantErr: nil,
		},
		{
			name: "発行日が支払日よりも後",
			args: args{
				issueDate: time.Now().AddDate(0, 0, 1),
				dueDate:   time.Now(),
				invoiceId: "invoice1",
				companyId: "company1",
				clientId:  "client1",
				payment:   int(paymentAmount),
			},
			want:    Invoice{},
			wantErr: ErrorIssueDateInvalid,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			invoice, err := NewInvoice(tt.args.issueDate, tt.args.dueDate, tt.args.invoiceId, tt.args.companyId, tt.args.clientId, tt.args.payment)
			assert.Equal(t, tt.want, invoice)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestInvoice_CalculateInvoiceAmount(t *testing.T) {
	issueDate := time.Now()
	dueDate := time.Now().AddDate(0, 1, 0)
	paymentAmount := 10000.00
	fee := paymentAmount * FEE_RATE
	tax := (paymentAmount * FEE_RATE) * TAX_RATE
	invoiceAmount := paymentAmount + (paymentAmount * FEE_RATE) + ((paymentAmount * FEE_RATE) * TAX_RATE)

	tests := []struct {
		name           string
		invoice        Invoice
		wantInvoiceAmt InvoiceAmount
		wantFee        Fee
		wantTax        Tax
		wantErr        error
	}{
		{
			name: "正常な計算",
			invoice: func() Invoice {
				i, _ := NewInvoice(issueDate, dueDate, "invoice1", "company1", "client1", int(paymentAmount))
				return i
			}(),
			wantInvoiceAmt: InvoiceAmount{value: invoiceAmount},
			wantFee:        Fee{value: fee},
			wantTax:        Tax{value: tax},
			wantErr:        nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.invoice.CalculateInvoiceAmount()
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantInvoiceAmt, tt.invoice.InvoiceAmount)
			assert.Equal(t, tt.wantFee, tt.invoice.Fee)
			assert.Equal(t, tt.wantTax, tt.invoice.Tax)
		})
	}
}
