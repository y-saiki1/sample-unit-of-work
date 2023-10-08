package domain

import "time"

type ComparableDate interface {
	IsAfter(date ComparableDate) bool
	IsBefore(date ComparableDate) bool
	Value() time.Time
}
