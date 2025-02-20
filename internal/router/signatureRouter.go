package router

import (
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/di"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/handler"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/handler/signature"
	"github.com/gorilla/mux"
)

func SignatureRouter(r *mux.Router, di di.DI) {
	// Create Struct Handler
	signatureHandler := signature.SignatureHandler{
		DI: di,
	}

	// Signature Route Handler
	r.HandleFunc("/signature", handler.ConvertToStandartHandlerFunc(signatureHandler.GetAllData)).Methods("GET")
	r.HandleFunc("/signature", handler.ConvertToStandartHandlerFunc(signatureHandler.CreateSignature)).Methods("POST")
	r.HandleFunc("/signature/{id}", handler.ConvertToStandartHandlerFunc(signatureHandler.UpdateSignature)).Methods("PATCH")
	r.HandleFunc("/signature/{id}", handler.ConvertToStandartHandlerFunc(signatureHandler.DeleteSignature)).Methods("DELETE")
}
