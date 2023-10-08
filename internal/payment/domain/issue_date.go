package domain

import (
	"time"
)

type IssueDate struct {
	value time.Time
}

func NewIssueDate(value time.Time) (IssueDate, error) {
	now := time.Now()
	day := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	startOfDay := day.Add(-1 * time.Nanosecond)
	endOfDay := day.AddDate(0, 0, 1)
	if !value.After(startOfDay) || !value.Before(endOfDay) {
		return IssueDate{}, ErrorIssueDateInvalid
	}
	return IssueDate{value}, nil
}

func (v IssueDate) IsAfter(date ComparableDate) bool {
	return v.value.After(date.Value())
}

func (v IssueDate) IsBefore(date ComparableDate) bool {
	return v.value.Before(date.Value())
}

func (v IssueDate) Value() time.Time {
	return v.value
}
