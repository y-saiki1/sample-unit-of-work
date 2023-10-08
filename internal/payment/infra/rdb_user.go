package infra

import (
	"upsidr-coding-test/internal/payment/domain"
	"upsidr-coding-test/internal/rdb"
)

type UserRDB struct{}

func NewUserRDB() UserRDB {
	return UserRDB{}
}

func (UserRDB) FindByUserId(userId domain.UserId) (domain.User, error) {
	user := rdb.User{}
	if err := rdb.DB.Where("user_id = ?", userId.Value()).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return userConverter{}.ToEntity(user)
}
