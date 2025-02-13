package router

import (
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/di"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/handler"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/handler/users"
	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router, di di.DI) {
	// Handler Users
	UserHandler := users.NewUsersHandler(di)

	// Routing Users
	r.HandleFunc("/users", handler.ConvertToStandartHandlerFunc(UserHandler.GetUsers)).Methods("GET")
	r.HandleFunc("/users", handler.ConvertToStandartHandlerFunc(UserHandler.CreateUsers)).Methods("POST")
	r.HandleFunc("/users/{id}", handler.ConvertToStandartHandlerFunc(UserHandler.DeleteUsers)).Methods("DELETE")
	r.HandleFunc("/users/{id}", handler.ConvertToStandartHandlerFunc(UserHandler.UpdateUsers)).Methods("PATCH")
}
