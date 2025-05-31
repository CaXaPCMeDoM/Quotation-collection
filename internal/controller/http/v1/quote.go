package v1

import (
	"citatnik/internal/controller/http/httperror"
	"citatnik/internal/controller/http/v1/dto/request"
	"citatnik/internal/controller/http/v1/dto/response"
	"citatnik/internal/entity"
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func (h *V1) AddQuote(w http.ResponseWriter, r *http.Request) error {
	var q request.Quote
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		return httperror.NewHTTPError(http.StatusBadRequest, "Can't parse the request body. Check the fields.", err)
	}

	ctx := r.Context()
	quote, err := h.q.Add(ctx, q.ToEntity())
	if err != nil {
		return httperror.NewHTTPError(http.StatusInternalServerError, "Can't added the quoter", err)
	}

	var resp response.Quote

	resp.ToResp(quote)

	return writeJSON(w, resp)
}

func (h *V1) GetQuotes(w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(r.Context(), time.Minute)
	defer cancel()

	author := r.URL.Query().Get("author")
	if author != "" {
		return h.GetQuotesByAuthor(w, r)
	}
	//Else we need return all

	results, err := h.q.GetAll(ctx)

	if err != nil {
		return httperror.NewHTTPError(
			http.StatusInternalServerError,
			"Error when receiving all quotes",
			err,
		)
	}

	resps := sliceEntitiesToResponse(results)

	return writeJSON(w, resps)
}

func (h *V1) GetQuotesByAuthor(w http.ResponseWriter, r *http.Request) error {
	author := r.URL.Query().Get("author")

	ctx := r.Context()

	if author != "" {
		results, err := h.q.GetByAuthor(ctx, author)

		if err != nil {
			return httperror.NewHTTPError(
				http.StatusInternalServerError,
				"There is no way to get an entity by author",
				err,
			)
		}

		resps := sliceEntitiesToResponse(results)

		return writeJSON(w, resps)
	}

	return httperror.NewHTTPError(
		http.StatusBadRequest,
		"Author in request is empty",
		errors.New("author is empty "),
	)
}

func (h *V1) GetRandomQuotes(w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(r.Context(), time.Minute)
	defer cancel()

	result, err := h.q.GetRandom(ctx)

	if err != nil {
		if errors.Is(err, entity.ErrMissingQuotes) {
			return httperror.NewHTTPError(http.StatusNotFound, "There are no quotes", err)
		} else {
			return httperror.NewHTTPError(http.StatusInternalServerError, "Error when searching for a random quote", err)
		}
	}

	var resp response.Quote

	resp.ToResp(result)

	return writeJSON(w, resp)
}

func (h *V1) DeleteQuoteByID(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]
	ctx := r.Context()

	if idStr == "" {
		return httperror.NewHTTPError(http.StatusBadRequest, "Empty id", errors.New("v1 - DeleteQuoteByID - idStr"))
	}

	err := h.q.DeleteByID(ctx, idStr)

	if err != nil {
		if errors.Is(err, entity.ErrNotFoundQuote) {
			return httperror.NewHTTPError(http.StatusNotFound, "quote not found", err)
		} else {
			return httperror.NewHTTPError(http.StatusInternalServerError, "Error when deleting a quote", err)
		}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func sliceEntitiesToResponse(results []entity.Quote) []response.Quote {

	resps := make([]response.Quote, 0, len(results))

	for _, val := range results {
		var resp response.Quote
		resp.ToResp(val)
		resps = append(resps, resp)
	}

	return resps
}

func writeJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return httperror.NewHTTPError(http.StatusInternalServerError, "Error with encoder result", err)
	}

	return nil
}
