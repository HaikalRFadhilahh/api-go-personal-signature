package userresponse

import (
	requesthandler "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/request"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/models"
)

type GetUsersAllResponse struct {
	StatusCode int                       `json:"statusCode"`
	Status     string                    `json:"status"`
	Message    string                    `json:"message"`
	Data       []models.UsersModel       `json:"data"`
	Pagination requesthandler.Pagination `json:"pagination"`
}
