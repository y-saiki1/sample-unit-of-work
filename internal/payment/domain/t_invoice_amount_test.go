package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInvoiceAmount(t *testing.T) {

	type args struct {
		paymentAmount PaymentAmount
		fee           Fee
		tax           Tax
	}
	tests := []struct {
		name    string
		args    args
		want    InvoiceAmount
		wantErr error
	}{
		{
			name: "正常な値での請求金額計算",
			args: args{
				paymentAmount: PaymentAmount{value: 10000},
				fee:           Fee{value: 10000 * FEE_RATE},
				tax:           Tax{value: (10000 * FEE_RATE) * TAX_RATE},
			},
			want:    InvoiceAmount{value: 10000 + (10000 * FEE_RATE) + ((10000 * FEE_RATE) * TAX_RATE)},
			wantErr: nil,
		},
		{
			name: "支払い金額がゼロ",
			args: args{
				paymentAmount: PaymentAmount{value: 0},
				fee:           Fee{value: 400},
				tax:           Tax{value: 40},
			},
			want:    InvoiceAmount{},
			wantErr: ErrorNegativeInvoiceAmount,
		},
		{
			name: "手数料がゼロ",
			args: args{
				paymentAmount: PaymentAmount{value: 10000},
				fee:           Fee{value: 0},
				tax:           Tax{value: 40},
			},
			want:    InvoiceAmount{},
			wantErr: ErrorNegativeInvoiceAmount,
		},
		{
			name: "消費税がゼロ",
			args: args{
				paymentAmount: PaymentAmount{value: 10000},
				fee:           Fee{value: 400},
				tax:           Tax{value: 0},
			},
			want:    InvoiceAmount{},
			wantErr: ErrorNegativeInvoiceAmount,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewInvoiceAmount(tt.args.paymentAmount, tt.args.fee, tt.args.tax)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInvoiceAmount_Value(t *testing.T) {
	tests := []struct {
		name string
		obj  InvoiceAmount
		want float64
	}{
		{
			name: "請求金額を正しく取得",
			obj:  InvoiceAmount{10440},
			want: 10440,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.obj.Value())
		})
	}
}
