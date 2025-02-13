package users

import (
	"encoding/json"
	"net/http"

	requesthandler "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/request"
	userresponse "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/response/users"
	usecaseusers "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/usecase/users"
)

func (u *UsersHandler) GetUsers(w http.ResponseWriter, r *http.Request) error {
	// Call Pagination
	pagination := requesthandler.InitPagination(r)
	search := r.URL.Query().Get("search")

	// Call UseCase
	res, err := usecaseusers.GetAllUsers(u.DI, pagination, pagination, search)
	if err != nil {
		return err
	}

	// Create Response
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(userresponse.GetUsersAllResponse{
		StatusCode: http.StatusOK,
		Status:     "success",
		Message:    "Data Users",
		Data:       res,
		Pagination: *pagination,
	})
}
