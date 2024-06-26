package rdb

import "time"

type Company struct {
	CompanyId          string     `json:"company_id" gorm:"primaryKey"`
	Name               string     `json:"name"`
	RepresentativeName string     `json:"representative_name"`
	PhoneNumber        string     `json:"phone_number"`
	PostalCode         string     `json:"postal_code"`
	Address            string     `json:"address"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at"`
}
