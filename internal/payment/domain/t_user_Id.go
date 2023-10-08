package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserId(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    UserId
		wantErr error
	}{
		{
			name:    "Valid ID",
			args:    args{id: "sampleID1"},
			want:    UserId{value: "sampleID1"},
			wantErr: nil,
		},
		{
			name:    "Empty ID",
			args:    args{id: ""},
			want:    UserId{},
			wantErr: ErrorUserIdEmpty,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewUserId(tt.args.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
