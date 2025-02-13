package users

import (
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/di"
)

type UsersHandler struct {
	DI di.DI
}

func NewUsersHandler(di di.DI) *UsersHandler {
	return &UsersHandler{
		DI: di,
	}
}
