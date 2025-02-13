package models

import "time"

type UsersModel struct {
	ID         int64             `json:"id"`
	Firstname  string            `json:"firstName"`
	MiddleName *string           `json:"middleName"`
	LastName   *string           `json:"lastName"`
	Signature  *[]SignatureModel `json:"signature"`
	CreatedAt  *time.Time        `json:"createAt"`
	UpdatedAt  *time.Time        `json:"updatedAt"`
}
