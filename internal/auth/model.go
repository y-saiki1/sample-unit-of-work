package auth

type User struct {
	Id       string
	Email    string
	Password string
}

func NewUser(id, email, hashedPassword string) (*User, error) {
	return &User{
		Id:       id,
		Email:    email,
		Password: hashedPassword,
	}, nil
}
