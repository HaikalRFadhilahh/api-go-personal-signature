package models

import "time"

type SignatureModel struct {
	ID         int64      `json:"id"`
	UserID     int64      `json:"userId"`
	Token      string     `json:"token"`
	Purpose    string     `json:"purpose"`
	ValidUntil string     `json:"validUntil"`
	User       UsersModel `json:"users"`
	CreatedAt  *time.Time `json:"createAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}
