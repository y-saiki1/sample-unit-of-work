package service

import "errors"

var (
	ErrorClientNotRelatedWithCompany      error = errors.New("client is not related with the company")
	ErrorClientRelationVerificationFailed error = errors.New("an error occurred while verifying the client relationship. please try again.")
)
