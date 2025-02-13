package userresponse

import usersrequest "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/request/users"

type UpdateUsersResponse struct {
	StatusCode int                             `json:"statusCode"`
	Status     string                          `json:"status"`
	Message    string                          `json:"message"`
	Data       usersrequest.UpdateUsersRequest `json:"data"`
}
