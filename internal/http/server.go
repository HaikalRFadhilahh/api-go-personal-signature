package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/config"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/di"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/middleware"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/pkg"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/router"
	"github.com/gorilla/mux"
)

/*
	Struct Credentials Server Host And Port Configuration
*/

type apiServer struct {
	serverConnectionString string
}

/*
	New API Server Connection
*/

func NewApiServer() *apiServer {
	return &apiServer{
		serverConnectionString: fmt.Sprintf("%v:%v", pkg.GetEnvValue("HOST", "127.0.0.1"), pkg.GetEnvValue("PORT", "8000")),
	}
}

/*
RUNNING SERVER CONNECTION
*/
func (s *apiServer) Run() {
	// Declare Mux Router
	r := mux.NewRouter()

	// Create Database Connection Init
	db := config.NewDB().InitDB()
	defer config.CloseConnection(db)

	// Dependency Injection Struct
	dependencyInjection := di.InitDependencyInjection(db)

	// Midleware
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.PanicHandler)

	// Routes Handler
	router.UserRouter(r, dependencyInjection)

	// Not Found Handler
	r.NotFoundHandler = router.NotFoundRoute()

	// Start Server
	fmt.Println("Server Running On ", s.serverConnectionString)
	if err := http.ListenAndServe(s.serverConnectionString, r); err != nil {
		log.Fatalln("Server Stopping!, Error :", err.Error())
	}
}
