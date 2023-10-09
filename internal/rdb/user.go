package rdb

import "time"

type User struct {
	UserId    string     `json:"user_id" gorm:"primaryKey"`
	CompanyId string     `json:"company_id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
