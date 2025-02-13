package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/apierror"
	usersrequest "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/request/users"
	userresponse "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/response/users"
	usecaseusers "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/usecase/users"
	"github.com/gorilla/mux"
)

func (u *UsersHandler) UpdateUsers(w http.ResponseWriter, r *http.Request) error {
	// ID Indentifier
	ID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return apierror.BadRequestError
	}

	// Struct Handler
	var req usersrequest.UpdateUsersRequest

	// usecase
	if err := usecaseusers.UpdateUsers(u.DI, &req, &ID, r); err != nil {
		return err
	}

	// Return Response
	return json.NewEncoder(w).Encode(&userresponse.UpdateUsersResponse{
		StatusCode: http.StatusOK,
		Status:     "success",
		Message:    "Update Data Users Success",
		Data:       req,
	})
}
