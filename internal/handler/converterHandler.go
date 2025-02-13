package handler

import (
	"net/http"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/apierror"
)

type handler func(http.ResponseWriter, *http.Request) error

func ConvertToStandartHandlerFunc(f handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			apierror.ReturnErrorResponse(w, r, err)
		}
	}

}
