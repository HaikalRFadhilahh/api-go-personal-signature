package router

import (
	"encoding/json"
	"net/http"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/apierror"
)

func NotFoundRoute() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(apierror.NotFoundError.StatusCode)
		json.NewEncoder(w).Encode(apierror.NotFoundError)
	})
}
