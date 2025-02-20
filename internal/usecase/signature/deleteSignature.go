package usecasesignature

import (
	"net/http"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/apierror"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/di"
)

func DeleteSignature(di di.DI, id int) error {
	// Declare Query
	query := "DELETE FROM signature WHERE id=?"

	// Exec Query
	r, e := di.GetDB().Exec(query, id)
	if e != nil {
		return apierror.ApiError{StatusCode: http.StatusInternalServerError, Status: "error", Message: e.Error()}
	}

	// Check Signature Users Exist or Not
	if c, _ := r.RowsAffected(); c <= 0 {
		return apierror.ApiError{StatusCode: http.StatusNotFound, Status: "error", Message: "Data Signature Not Found!"}
	}

	return nil
}
