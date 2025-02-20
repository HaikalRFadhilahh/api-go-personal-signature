package signatureresponse

import signaturerequest "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/request/signature"

type UpdateSignatureResponse struct {
	StatusCode int                                     `json:"statusCode"`
	Status     string                                  `json:"status"`
	Message    string                                  `json:"message"`
	Data       signaturerequest.UpdateSignatureRequest `json:"data"`
}
