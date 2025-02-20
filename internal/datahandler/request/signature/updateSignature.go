package signaturerequest

type UpdateSignatureRequest struct {
	UserID     int    `json:"userId"`
	Purpose    string `json:"purpose"`
	ValidUntil string `json:"validUntil"`
}
