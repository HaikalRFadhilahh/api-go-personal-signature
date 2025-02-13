package usecaseusers

import (
	"net/http"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/apierror"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/di"
)

func DeleteUsers(di di.DI, id *int) error {
	// Create Query Delete
	queryDelete := "DELETE FROM users WHERE id=?"

	// Exec Query Delete
	r, err := di.GetDB().Exec(queryDelete, id)
	if err != nil {
		return apierror.ApiError{StatusCode: http.StatusInternalServerError, Status: "error", Message: err.Error()}
	}

	if c, _ := r.RowsAffected(); c < 1 {
		return apierror.ApiError{StatusCode: http.StatusNotFound, Status: "error", Message: "Users Data Not Found!"}
	}

	return nil
}
