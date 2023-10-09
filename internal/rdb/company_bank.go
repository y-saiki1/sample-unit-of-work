package rdb

import "time"

type CompanyBank struct {
	CompanyBankId string     `json:"company_bank_id" gorm:"primaryKey"`
	CompanyId     string     `json:"company_id"`
	BankCode      string     `json:"bank_code"`
	BranchCode    string     `json:"branch_code"`
	BankName      string     `json:"bank_name"`
	BranchName    string     `json:"branch_name"`
	AccountType   string     `json:"account_type"`
	AccountNumber string     `json:"account_number"`
	AccountHolder string     `json:"account_holder"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}
