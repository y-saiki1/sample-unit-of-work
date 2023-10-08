package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInvoiceId(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    InvoiceId
		wantErr error
	}{
		{
			name:    "Valid ID",
			args:    args{id: "sampleID1"},
			want:    InvoiceId{value: "sampleID1"},
			wantErr: nil,
		},
		{
			name:    "Empty ID",
			args:    args{id: ""},
			want:    InvoiceId{},
			wantErr: ErrorInvoiceIdEmpty,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewInvoiceId(tt.args.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
