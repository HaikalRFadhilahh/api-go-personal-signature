package signature

import (
	"encoding/json"
	"net/http"

	requesthandler "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/request"
	signatureresponse "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/response/signature"
	usecasesignature "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/usecase/signature"
)

func (s *SignatureHandler) GetAllData(w http.ResponseWriter, r *http.Request) error {
	// Get Search Query
	search := r.URL.Query().Get("search")

	// Init Pagination Handler
	pagination := requesthandler.InitPagination(r)

	// UseCase Call Get All Signature Data
	data, err := usecasesignature.GetAllSignature(s.DI, search, pagination)
	if err != nil {
		return err
	}

	// Return json
	return json.NewEncoder(w).Encode(&signatureresponse.GetAllSignatureResponse{
		StatusCode: http.StatusOK,
		Status:     "success",
		Message:    "Data Signature With Relation",
		Data:       data,
		Pagination: *pagination,
	})
}
