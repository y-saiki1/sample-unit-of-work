package domain

const (
	INVOICE_STATUS_PENDING    = "未処理"
	INVOICE_STATUS_PROCESSING = "処理中"
	INVOICE_STATUS_PAID       = "支払い済み"
	INVOICE_STATUS_ERROR      = "エラー"
)

type InvoiceStatus struct {
	value string
}

func NewInvoiceStatus(status string) (InvoiceStatus, error) {
	switch status {
	case INVOICE_STATUS_PENDING, INVOICE_STATUS_PROCESSING, INVOICE_STATUS_PAID, INVOICE_STATUS_ERROR:
		return InvoiceStatus{value: status}, nil
	}
	return InvoiceStatus{}, ErrorInvoiceStatusInvalid
}
