package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDueDate(t *testing.T) {
	value := time.Now()
	before := value.AddDate(0, 0, -1)
	after := value.AddDate(0, 0, 1)

	type args struct {
		value     time.Time
		issueDate IssueDate
	}
	tests := []struct {
		name    string
		args    args
		want    DueDate
		wantErr error
	}{
		{
			name: "発行日よりも後の期日",
			args: args{
				value:     value,
				issueDate: IssueDate{value: before},
			},
			want:    DueDate{value: value},
			wantErr: nil,
		},
		{
			name: "発行日よりも前の期日",
			args: args{
				value:     value,
				issueDate: IssueDate{value: after},
			},
			want:    DueDate{},
			wantErr: ErrorDueDateBeforeIssue,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewDueDate(tt.args.value, tt.args.issueDate)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDueDate_IsAfter(t *testing.T) {
	now := time.Now()
	dueDate := DueDate{value: now}

	tests := []struct {
		name    string
		duedate DueDate
		arg     ComparableDate
		want    bool
	}{
		{
			name: "引数の日付よりも後",
			arg:  IssueDate{value: now.AddDate(0, 0, -1)},
			want: true,
		},
		{
			name: "引数の日付と同じ",
			arg:  IssueDate{value: now},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := dueDate.IsAfter(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDueDate_IsBefore(t *testing.T) {
	now := time.Now()
	dueDate := DueDate{value: now}
	tests := []struct {
		name string
		arg  ComparableDate
		want bool
	}{
		{
			name: "引数の日付よりも前",
			arg:  IssueDate{value: now.AddDate(0, 0, 1)},
			want: true,
		},
		{
			name: "引数の日付と同じ",
			arg:  IssueDate{value: now},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := dueDate.IsBefore(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDueDate_Value(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name string
		due  DueDate
		want time.Time
	}{
		{
			name: "取得した発行日の確認",
			due:  DueDate{value: now},
			want: now,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.due.Value())
		})
	}
}
