package domain

type CompanyRepository interface {
	IsClientOfCompany(companyId CompanyId, clientId ClientId) (bool, error)
}

type UserRepository interface {
	FindByUserId(userId UserId) (User, error)
}
