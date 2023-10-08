package domain

import (
	"time"
)

type DueDate struct {
	value time.Time
}

func NewDueDate(value time.Time, issueDate IssueDate) (DueDate, error) {
	dueDate := DueDate{value}
	if dueDate.IsBefore(issueDate) {
		return DueDate{}, ErrorDueDateBeforeIssue
	}
	return dueDate, nil
}

func (v DueDate) IsAfter(date ComparableDate) bool {
	return v.value.After(date.Value())
}

func (v DueDate) IsBefore(date ComparableDate) bool {
	return v.value.Before(date.Value())
}

func (v DueDate) Value() time.Time {
	return v.value
}
