package domain

const (
	FEE_RATE = 0.04
)

type Fee struct {
	value float64
}

func NewFee(paymentAmount PaymentAmount) Fee {
	fee := paymentAmount.Value() * FEE_RATE
	return Fee{fee}
}

func (v Fee) Value() float64 {
	return v.value
}
