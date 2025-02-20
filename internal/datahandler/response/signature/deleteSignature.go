package signatureresponse

type DeleteSignatureResponse struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
