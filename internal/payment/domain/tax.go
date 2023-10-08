package domain

const (
	TAX_RATE = 0.10
)

type Tax struct {
	value float64
}

func NewTax(fee Fee) Tax {
	tax := fee.Value() * TAX_RATE
	return Tax{tax}
}

func (v Tax) Value() float64 {
	return v.value
}
