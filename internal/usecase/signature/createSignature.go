package usecasesignature

import (
	"net/http"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/apierror"
	signaturerequest "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/request/signature"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/di"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/pkg"
)

func CreateSignature(di di.DI, req *signaturerequest.CreateSignatureRequest) error {
	// Check Users ID
	q := "SELECT COUNT(*) FROM users WHERE id=?"
	if err := di.GetDB().QueryRow(q, req.UserID).Err(); err != nil {
		return apierror.ApiError{StatusCode: http.StatusBadRequest, Status: "error", Message: "User ID Not Found!"}
	}

	// Generate Random Token
	randstr, err := pkg.GenerateRandomString(20)
	if err != nil {
		return apierror.ApiError{StatusCode: http.StatusInternalServerError, Status: "error", Message: err.Error()}
	}
	req.Token = randstr

	// Query
	q = "INSERT INTO signature(user_id,token,purpose,valid_until) VALUES (?,?,?,?)"
	if _, err := di.GetDB().Exec(q, req.UserID, req.Token, req.Purpose, req.ValidUntil); err != nil {
		return apierror.ApiError{StatusCode: http.StatusInternalServerError, Status: "error", Message: err.Error()}
	}

	return nil
}
