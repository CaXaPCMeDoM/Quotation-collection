package v1

import (
	"citatnik/internal/controller/http/httperror"
	"citatnik/internal/usecase"
	"github.com/gorilla/mux"
	"net/http"
)

func NewQuoteRoutes(
	r *mux.Router,
	q usecase.Quote,
) {
	handler := &V1{q: q}

	sr := r.PathPrefix("/quotes").Subrouter()
	sr.HandleFunc("", httperror.WrapNetHTTP(handler.AddQuote)).Methods(http.MethodPost)
	sr.HandleFunc("", httperror.WrapNetHTTP(handler.GetQuotes)).Methods(http.MethodGet)
	sr.HandleFunc("/random", httperror.WrapNetHTTP(handler.GetRandomQuotes)).Methods(http.MethodGet)
	sr.HandleFunc("/{id}", httperror.WrapNetHTTP(handler.DeleteQuoteByID)).Methods(http.MethodDelete)
}
