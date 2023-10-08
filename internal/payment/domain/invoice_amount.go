package domain

type InvoiceAmount struct {
	value float64
}

func NewInvoiceAmount(paymentAmount PaymentAmount, fee Fee, tax Tax) (InvoiceAmount, error) {
	if paymentAmount.Value() <= 0 || fee.Value() <= 0 || tax.Value() <= 0 {
		return InvoiceAmount{}, ErrorNegativeInvoiceAmount
	}
	total := paymentAmount.Value() + fee.Value() + tax.Value()
	return InvoiceAmount{total}, nil
}

func (v InvoiceAmount) Value() float64 {
	return v.value
}
