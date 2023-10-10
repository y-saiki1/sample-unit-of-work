package domain

type InvoiceStatus string

func (v InvoiceStatus) Value() string {
	return string(v)
}

const (
	INVOICE_STATUS_PENDING    InvoiceStatus = "未処理"
	INVOICE_STATUS_PROCESSING InvoiceStatus = "処理中"
	INVOICE_STATUS_PAID       InvoiceStatus = "支払い済み"
	INVOICE_STATUS_ERROR      InvoiceStatus = "エラー"
)

type InvoiceStatusLog struct {
	UserId   UserId
	UserName UserName
	Status   InvoiceStatus
}

func NewInvoiceStatusLog(userId, userName string, status InvoiceStatus) (InvoiceStatusLog, error) {
	uId, err := NewUserId(userId)
	if err != nil {
		return InvoiceStatusLog{}, ErrorUserIdEmpty
	}
	uName, err := NewUserName(userName)
	if err != nil {
		return InvoiceStatusLog{}, ErrorUserNameEmpty
	}

	switch status {
	case INVOICE_STATUS_PENDING, INVOICE_STATUS_PROCESSING, INVOICE_STATUS_PAID, INVOICE_STATUS_ERROR:
		return InvoiceStatusLog{Status: InvoiceStatus(status), UserId: uId, UserName: uName}, nil
	}
	return InvoiceStatusLog{}, ErrorInvoiceStatusInvalid
}
