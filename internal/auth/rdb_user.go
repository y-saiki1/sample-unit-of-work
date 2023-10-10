package auth

import "upsidr-coding-test/internal/rdb"

type UserRDB struct{}

func NewUserRDB() UserRDB {
	return UserRDB{}
}

func (UserRDB) FindByEmail(email string) (User, error) {
	u := rdb.User{}
	if err := rdb.DB.
		Where("email = ?", email).
		First(&u).Error; err != nil {
		return User{}, nil
	}

	return User{
		Id:       u.UserId,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}
