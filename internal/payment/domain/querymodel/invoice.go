package querymodel

import "time"

type InvoiceQueryFilter struct {
	Offset          int        `json:"offset"`
	Limit           int        `json:"limit"`
	CompanyId       string     `json:"companyId"`
	ContainsDeleted bool       `json:"excludeDeleted"`
	DueFrom         *time.Time `json:"dueFrom"`
	DueTo           *time.Time `json:"dueTo"`
}

type InvoiceQueryModel struct {
	InvoiceId     string     `json:"invoiceId"`
	CompanyId     string     `json:"companyId"`
	PaymentAmount float64    `json:"paymentAmount"`
	Fee           float64    `json:"fee"`
	FeeRate       float64    `json:"feeRate"`
	Tax           float64    `json:"tax"`
	TaxRate       float64    `json:"taxRate"`
	InvoiceAmount float64    `json:"invoiceAmount"`
	DueDate       *time.Time `json:"dueDate"`
	IssueDate     *time.Time `json:"issueDate"`
	UpdatedAt     *time.Time `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt"`

	Client     InvoiceClientQueryModel `json:"client"`
	StatusLogs []InvoiceStatusLogQueryModel
}

type InvoiceStatusLogQueryModel struct {
	UserName  string     `json:"userName"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type InvoiceClientQueryModel struct {
	CompanyId string `json:"companyId"`
	Name      string `json:"name"`
}
