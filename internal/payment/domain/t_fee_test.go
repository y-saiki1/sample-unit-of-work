package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFee(t *testing.T) {
	type args struct {
		paymentAmount PaymentAmount
	}
	tests := []struct {
		name string
		args args
		want Fee
	}{
		{
			name: "手数料の計算（10000円の場合）",
			args: args{paymentAmount: PaymentAmount{value: 10000}},
			want: Fee{value: 10000 * FEE_RATE},
		},
		{
			name: "手数料の計算（5000円の場合）",
			args: args{PaymentAmount{value: 5000}},
			want: Fee{value: 5000 * FEE_RATE},
		},
		{
			name: "手数料の計算（0円の場合）",
			args: args{PaymentAmount{value: 0}},
			want: Fee{value: 0 * FEE_RATE},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewFee(tt.args.paymentAmount)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFee_Value(t *testing.T) {
	tests := []struct {
		name string
		fee  Fee
		want float64
	}{
		{
			name: "手数料の取得（40円の場合）",
			fee:  Fee{value: 40},
			want: 40,
		},
		{
			name: "手数料の取得（200円の場合）",
			fee:  Fee{value: 200},
			want: 200,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.fee.Value()
			assert.Equal(t, tt.want, got)
		})
	}
}
