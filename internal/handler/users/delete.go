package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/apierror"
	userresponse "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/response/users"
	usecaseusers "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/usecase/users"
	"github.com/gorilla/mux"
)

func (u *UsersHandler) DeleteUsers(w http.ResponseWriter, r *http.Request) error {
	// Get Params Data Id Users
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return apierror.BadRequestError
	}

	// Exec Use Case For Remove Data
	if err := usecaseusers.DeleteUsers(u.DI, &id); err != nil {
		return err
	}

	// Return Json Response
	return json.NewEncoder(w).Encode(userresponse.DeleteUsersResponse{StatusCode: http.StatusOK, Status: "success", Message: "Delete Users Success", Data: map[string]int{"id": id}})
}
