package domain

type CompanyId struct {
	value string
}

func NewCompanyId(id string) (CompanyId, error) {
	if id == "" {
		return CompanyId{}, ErrorCompanyIdEmpty
	}
	return CompanyId{value: id}, nil
}

func (v CompanyId) Value() string {
	return v.value
}
