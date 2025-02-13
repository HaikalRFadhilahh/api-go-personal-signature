package users

import (
	"encoding/json"
	"net/http"

	apierror "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/apierror"
	usersrequest "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/request/users"
	userresponse "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/response/users"
	usecaseusers "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/usecase/users"
)

func (u *UsersHandler) CreateUsers(w http.ResponseWriter, r *http.Request) error {
	// Bind Response
	requestData := usersrequest.CreateUsersRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		return apierror.ApiError{StatusCode: 403, Status: "error", Message: err.Error()}
	}

	// Insert into Database
	res, err := usecaseusers.CreateUsers(u.DI, &requestData)
	if err != nil {
		return err
	}

	// Return Response
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(userresponse.CreateUsersResponse{StatusCode: http.StatusOK, Status: "success", Message: "Success Create Data Users", Data: *res})
}
