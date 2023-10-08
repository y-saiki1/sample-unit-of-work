package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientId(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    ClientId
		wantErr error
	}{
		{
			name:    "Valid ID",
			args:    args{id: "sampleID1"},
			want:    ClientId{value: "sampleID1"},
			wantErr: nil,
		},
		{
			name:    "Empty ID",
			args:    args{id: ""},
			want:    ClientId{},
			wantErr: ErrorClientIdEmpty,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewClientId(tt.args.id)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
