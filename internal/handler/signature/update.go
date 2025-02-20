package signature

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/apierror"
	signaturerequest "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/request/signature"
	signatureresponse "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/response/signature"
	usecasesignature "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/usecase/signature"
	"github.com/gorilla/mux"
)

func (s *SignatureHandler) UpdateSignature(w http.ResponseWriter, r *http.Request) error {
	// Decode ID Params
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return apierror.ApiError{
			StatusCode: http.StatusBadRequest,
			Status:     "error",
			Message:    "ID Signature Params Data Not Found!",
		}
	}

	// Declare Request Variabel
	var req signaturerequest.UpdateSignatureRequest

	// UseCase Update Data Signature
	if err := usecasesignature.UpdateSignature(s.DI, r, &req, id); err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(&signatureresponse.UpdateSignatureResponse{
		StatusCode: http.StatusOK,
		Status:     "success",
		Message:    "Success Updated Signature",
		Data:       req,
	})
}
