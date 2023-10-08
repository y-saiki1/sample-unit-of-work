package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPaymentAmount(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name    string
		args    args
		want    PaymentAmount
		wantErr error
	}{
		{
			name:    "正の値",
			args:    args{value: 100},
			want:    PaymentAmount{value: 100.0},
			wantErr: nil,
		},
		{
			name:    "ゼロ",
			args:    args{value: 0},
			want:    PaymentAmount{value: 0},
			wantErr: ErrorInvalidPaymentAmount,
		},
		{
			name:    "負の値",
			args:    args{value: -100},
			want:    PaymentAmount{},
			wantErr: ErrorInvalidPaymentAmount,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewPaymentAmount(tt.args.value)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPaymentAmount_Value(t *testing.T) {
	tests := []struct {
		name    string
		arg     int
		want    float64
		wantErr error
	}{
		{
			name:    "正の値を取得",
			arg:     100,
			want:    100.0,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			paymentAmount, err := NewPaymentAmount(tt.arg)
			got := paymentAmount.Value()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
