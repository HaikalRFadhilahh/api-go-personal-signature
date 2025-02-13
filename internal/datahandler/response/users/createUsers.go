package userresponse

import (
	usersrequest "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/request/users"
)

type CreateUsersResponse struct {
	StatusCode int                             `json:"statusCode"`
	Status     string                          `json:"status"`
	Message    string                          `json:"message"`
	Data       usersrequest.CreateUsersRequest `json:"data"`
}
