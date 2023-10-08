package domain

type UnitOfWorkExecutor interface {
	Exec() error
}

type InvoiceRepository interface {
	Store(invoice Invoice) error
}

type CompanyRepository interface {
	IsClientOfCompany(companyId CompanyId, clientId ClientId) (bool, error)
}

type UserRepository interface {
	FindByUserId(userId UserId) (User, error)
}
