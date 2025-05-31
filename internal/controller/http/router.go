package http

import (
	"citatnik/internal/controller/http/middleware"
	v1 "citatnik/internal/controller/http/v1"
	"citatnik/internal/usecase"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(quote usecase.Quote) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware.Recovery())

	apiV1 := r.PathPrefix("").Subrouter()
	v1.NewQuoteRoutes(apiV1, quote)

	return r
}
