package usecasesignature

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/apierror"
	signaturerequest "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/request/signature"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/di"
)

func UpdateSignature(di di.DI, r *http.Request, req *signaturerequest.UpdateSignatureRequest, id int) error {
	// Query
	queryUpdate := "UPDATE signature SET user_id=?,purpose=?,valid_until=? WHERE id=?"
	querySelectCheck := "SELECT user_id,purpose,valid_until FROM signature WHERE id=?"
	queryUsersCheck := "SELECT COUNT(*) FROM users WHERE id=?"

	// Check Signature Data
	if err := di.GetDB().QueryRow(querySelectCheck, id).Scan(&req.UserID, &req.Purpose, &req.ValidUntil); err != nil {
		return apierror.ApiError{StatusCode: http.StatusNotFound, Status: "error", Message: "Data Signature Not Found!"}
	}

	// Binding Json Request
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return apierror.ApiError{StatusCode: http.StatusBadRequest, Status: "error", Message: err.Error()}
	}

	// Check Users ID
	var usersCount int64
	if err := di.GetDB().QueryRow(queryUsersCheck, &req.UserID).Scan(&usersCount); err != nil || usersCount <= 0 {
		return apierror.ApiError{StatusCode: http.StatusNotFound, Status: "error", Message: "Data Users ID Not Found!"}
	}

	// Exec
	if _, err := di.GetDB().Exec(queryUpdate, req.UserID, req.Purpose, req.ValidUntil, id); err != nil {
		fmt.Println(err.Error())
		return apierror.ApiError{StatusCode: http.StatusInternalServerError, Status: "error", Message: "Internal Server Error"}
	}

	return nil
}
