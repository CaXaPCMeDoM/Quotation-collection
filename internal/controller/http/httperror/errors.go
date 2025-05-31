package httperror

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type HandlerFuncWithError func(w http.ResponseWriter, r *http.Request) error

type HTTPError struct {
	Code       int
	Message    string
	InnerError error
}

func NewHTTPError(code int, message string, inner error) *HTTPError {
	return &HTTPError{
		Code:       code,
		Message:    message,
		InnerError: inner,
	}
}
func (e *HTTPError) Error() string {
	return e.Message
}

func WrapNetHTTP(endpoint HandlerFuncWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := endpoint(w, r); err != nil {
			var httpErr *HTTPError
			if errors.As(err, &httpErr) {
				if httpErr.InnerError != nil {
					log.Printf("Client Message: %s, Internal Error: %s. Status Code: %d", httpErr.Message, httpErr.InnerError, httpErr.Code)
				} else {
					log.Printf("HTTP error: %d %s", httpErr.Code, httpErr.Message)
				}
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(httpErr.Code)
				json.NewEncoder(w).Encode(map[string]string{"error": httpErr.Message})
			} else {
				log.Println("Internal server error:", err)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "Internal Server Error"})
			}
		}
	}
}
