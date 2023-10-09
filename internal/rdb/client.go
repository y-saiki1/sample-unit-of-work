package rdb

type Client struct {
	CompanyId string `json:"company_id" gorm:"primaryKey"`
	ClientId  string `json:"client_id" gorm:"primaryKey"`
}
