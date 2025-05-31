package entity

import "errors"

var (
	ErrMissingQuotes = errors.New("quotes are missing")
	ErrNotFoundQuote = errors.New("quote not found")
)
