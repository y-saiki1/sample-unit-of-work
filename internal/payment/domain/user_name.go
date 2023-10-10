package domain

type UserName struct {
	value string
}

func NewUserName(name string) (UserName, error) {
	if name == "" {
		return UserName{}, ErrorUserNameEmpty
	}
	return UserName{name}, nil
}

func (v UserName) Value() string {
	return v.value
}
