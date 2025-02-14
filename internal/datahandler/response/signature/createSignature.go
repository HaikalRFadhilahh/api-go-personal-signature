package signatureresponse

import signaturerequest "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/request/signature"

type CreateSignatureResponse struct {
	StatusCode int                                     `json:"statusCode"`
	Status     string                                  `json:"status"`
	Message    string                                  `json:"message"`
	Data       signaturerequest.CreateSignatureRequest `json:"data"`
}
