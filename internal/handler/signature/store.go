package signature

import (
	"encoding/json"
	"net/http"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/apierror"
	signaturerequest "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/request/signature"
	signatureresponse "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/response/signature"
	usecasesignature "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/usecase/signature"
)

func (s *SignatureHandler) CreateSignature(w http.ResponseWriter, r *http.Request) error {
	// Binding Request
	var req signaturerequest.CreateSignatureRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.ValidUntil == "" {
		return apierror.BadRequestError
	}

	// Usecase Handler
	if err := usecasesignature.CreateSignature(s.DI, &req); err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(&signatureresponse.CreateSignatureResponse{
		StatusCode: http.StatusOK,
		Status:     "success",
		Message:    "Success Add Data Signature",
		Data:       req,
	})
}
