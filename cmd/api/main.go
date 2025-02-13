package main

import (
	"fmt"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/http"
	"github.com/joho/godotenv"
)

/*
	LOAD Default External Dependencies
*/

func init() {
	// Load .ENV
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error Importing .env File, Error :", err.Error())
	}
}

func main() {
	// Starting Server
	http.NewApiServer().Run()
}
