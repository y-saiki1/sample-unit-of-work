package domain

type InvoiceId struct {
	value string
}

func NewInvoiceId(id string) (InvoiceId, error) {
	if id == "" {
		return InvoiceId{}, ErrorInvoiceIdEmpty
	}
	return InvoiceId{value: id}, nil
}

func (v InvoiceId) Value() string {
	return v.value
}
