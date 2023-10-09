package queryservice

import "time"

type InvoiceListDTO struct {
	UserId          string
	DueFrom         *time.Time
	DueTo           *time.Time
	Page            int
	ContainsDeleted bool
}
