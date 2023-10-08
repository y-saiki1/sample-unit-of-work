package domain

type PaymentAmount struct {
	value float64
}

func NewPaymentAmount(value int) (PaymentAmount, error) {
	if value <= 0 {
		return PaymentAmount{}, ErrorInvalidPaymentAmount
	}
	return PaymentAmount{float64(value)}, nil
}

func (v PaymentAmount) Value() float64 {
	return v.value
}
