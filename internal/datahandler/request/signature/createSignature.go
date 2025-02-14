package signaturerequest

type CreateSignatureRequest struct {
	UserID     int    `json:"userId"`
	Token      string `json:"-"`
	Purpose    string `json:"purpose"`
	ValidUntil string `json:"validUntil"`
}
