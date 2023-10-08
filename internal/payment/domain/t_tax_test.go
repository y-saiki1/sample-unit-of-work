package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTax(t *testing.T) {
	type args struct {
		fee Fee
	}
	tests := []struct {
		name string
		args args
		want Tax
	}{
		{
			name: "正常な手数料からの税金計算",
			args: args{fee: Fee{value: 1000.0}},
			want: Tax{value: 100.0},
		},
		{
			name: "手数料が0の場合の税金計算",
			args: args{fee: Fee{value: 0.0}},
			want: Tax{value: 0.0},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewTax(tt.args.fee)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTax_Value(t *testing.T) {
	tests := []struct {
		name string
		tax  Tax
		want float64
	}{
		{
			name: "正常な税金値の取得",
			tax:  Tax{value: 100.0},
			want: 100.0,
		},
		{
			name: "0の税金値の取得",
			tax:  Tax{value: 0.0},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.tax.Value()
			assert.Equal(t, tt.want, got)
		})
	}
}
