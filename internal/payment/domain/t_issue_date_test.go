package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewIssueDate(t *testing.T) {
	value := time.Now()
	before := value.AddDate(0, 0, -1)
	after := value.AddDate(0, 0, 1)

	type args struct {
		value time.Time
	}

	tests := []struct {
		name    string
		args    args
		want    IssueDate
		wantErr error
	}{
		{
			name: "現在時刻の発行日",
			args: args{
				value: value,
			},
			want:    IssueDate{value: value},
			wantErr: nil,
		},
		{
			name: "未来の発行日",
			args: args{
				value: after,
			},
			want:    IssueDate{},
			wantErr: ErrorIssueDateInvalid,
		},
		{
			name: "過去の発行日",
			args: args{
				value: before,
			},
			want:    IssueDate{},
			wantErr: ErrorIssueDateInvalid,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewIssueDate(tt.args.value)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIssueDate_IsAfter(t *testing.T) {
	now := time.Now()
	type args struct {
		date ComparableDate
	}
	tests := []struct {
		name string
		id   IssueDate
		args args
		want bool
	}{
		{
			name: "発行日は後",
			id:   IssueDate{value: now},
			args: args{
				date: IssueDate{value: now.Add(-1 * time.Hour)},
			},
			want: true,
		},
		{
			name: "発行日は前",
			id:   IssueDate{value: now},
			args: args{
				date: IssueDate{value: now.Add(1 * time.Hour)},
			},
			want: false,
		},
		{
			name: "同じ発行日",
			id:   IssueDate{value: now},
			args: args{
				date: IssueDate{value: now},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.id.IsAfter(tt.args.date))
		})
	}
}

func TestIssueDate_IsBefore(t *testing.T) {
	now := time.Now()
	issdate := IssueDate{value: now}

	type args struct {
		date ComparableDate
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "発行日は後",
			args: args{
				date: IssueDate{value: now.Add(1 * time.Hour)},
			},
			want: true,
		},
		{
			name: "発行日は前",
			args: args{
				date: IssueDate{value: now.Add(-1 * time.Hour)},
			},
			want: false,
		},
		{
			name: "同じ発行日",
			args: args{
				date: IssueDate{value: now},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, issdate.IsBefore(tt.args.date))
		})
	}
}

func TestIssueDate_Value(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name string
		id   IssueDate
		want time.Time
	}{
		{
			name: "取得した発行日の確認",
			id:   IssueDate{value: now},
			want: now,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.id.Value())
		})
	}
}
