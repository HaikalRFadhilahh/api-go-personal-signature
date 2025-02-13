package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/apierror"
)

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Panic Error :", err)
				w.WriteHeader(apierror.InternalServerError.StatusCode)
				json.NewEncoder(w).Encode(apierror.InternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
