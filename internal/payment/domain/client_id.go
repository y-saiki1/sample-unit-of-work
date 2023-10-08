package domain

type ClientId struct {
	value string
}

func NewClientId(id string) (ClientId, error) {
	if id == "" {
		return ClientId{}, ErrorClientIdEmpty
	}
	return ClientId{value: id}, nil
}

func (v ClientId) Value() string {
	return v.value
}
