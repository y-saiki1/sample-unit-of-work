package infra

import (
	"upsidr-coding-test/internal/payment/domain"
	"upsidr-coding-test/internal/rdb"

	"gorm.io/gorm"
)

type CompanyRDB struct {
}

func NewCompanyRDB() CompanyRDB {
	return CompanyRDB{}
}

func (r *CompanyRDB) IsClientOfCompany(companyId domain.CompanyId, clientId domain.ClientId) (bool, error) {
	client := struct {
		CompanyId string `json:"company_id"`
		ClientId  string `json:"client_id"`
	}{}
	if err := rdb.DB.
		Table("clients").
		Where("company_id = ?", companyId.Value()).
		Where("client_id = ?", clientId.Value()).
		First(&client).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
