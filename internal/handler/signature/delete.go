package signature

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/apierror"
	signatureresponse "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/response/signature"
	usecasesignature "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/usecase/signature"
	"github.com/gorilla/mux"
)

func (s *SignatureHandler) DeleteSignature(w http.ResponseWriter, r *http.Request) error {
	// Catch ID Signature Delete
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return apierror.ApiError{StatusCode: http.StatusBadRequest, Status: "error", Message: "Bad Request, ID Users Params Not Found!"}
	}

	// Delete Data Signature
	if err := usecasesignature.DeleteSignature(s.DI, id); err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(signatureresponse.DeleteSignatureResponse{
		StatusCode: http.StatusOK,
		Status:     "success",
		Message:    "Success Deleteed Data Signature",
		Data:       map[string]int{"id": id},
	})

}
