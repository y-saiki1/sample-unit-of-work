package domain

type UserId struct {
	value string
}

func NewUserId(id string) (UserId, error) {
	if id == "" {
		return UserId{}, ErrorUserIdEmpty
	}
	return UserId{value: id}, nil
}

func (v UserId) Value() string {
	return v.value
}
