package usecaseusers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/apierror"
	usersrequest "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/request/users"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/di"
)

func UpdateUsers(di di.DI, data *usersrequest.UpdateUsersRequest, id *int, r *http.Request) error {

	// Scan Old Data User
	q := "SELECT first_name,middle_name,last_name FROM users where id=?"
	if err := di.GetDB().QueryRow(q, id).Scan(&data.FirstName, &data.MiddleName, &data.LastName); err != nil {
		if err == sql.ErrNoRows {
			return apierror.ApiError{StatusCode: http.StatusNotFound, Status: "error", Message: "Data Users Not Found!"}
		}
		return apierror.ApiError{StatusCode: http.StatusInternalServerError, Status: "error", Message: err.Error()}
	}

	// Read Request Users
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return apierror.ApiError{StatusCode: http.StatusInternalServerError, Status: "error", Message: err.Error()}
	}

	// Update Users Data
	q = "UPDATE users SET first_name=?,middle_name=?,last_name=? WHERE id=?"
	if _, err := di.GetDB().Exec(q, data.FirstName, data.MiddleName, data.LastName, id); err != nil {
		return apierror.ApiError{StatusCode: http.StatusInternalServerError, Status: "error", Message: err.Error()}
	}

	return nil
}
