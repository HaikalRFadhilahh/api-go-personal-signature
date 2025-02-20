package apierror

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

func (a ApiError) Error() string {
	return a.Message
}

var (
	NotFoundError       = ApiError{404, "error", "not found"}
	InternalServerError = ApiError{500, "error", "internal server error"}
	BadRequestError     = ApiError{403, "error", "bad request"}
)

func ReturnErrorResponse(w http.ResponseWriter, r *http.Request, a any) {
	ta, ok := a.(ApiError)
	if !ok {
		w.WriteHeader(InternalServerError.StatusCode)
		json.NewEncoder(w).Encode(InternalServerError)
		return
	}
	w.WriteHeader(ta.StatusCode)
	json.NewEncoder(w).Encode(ta)
}
