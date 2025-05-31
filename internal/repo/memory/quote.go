package memory

import (
	"citatnik/internal/entity"
	"citatnik/internal/utils/counter"
	"context"
	"fmt"
	"math/rand"
	"sync"
)

type QuoteRepo struct {
	mu        sync.RWMutex
	quotes    map[string]entity.Quote
	idCurrent string
}

func New() *QuoteRepo {
	return &QuoteRepo{
		quotes:    make(map[string]entity.Quote),
		idCurrent: "0",
	}
}

func (r *QuoteRepo) Create(ctx context.Context, quote entity.Quote) (entity.Quote, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := counter.Increment(r.idCurrent)
	quote.ID = id
	r.idCurrent = id

	r.quotes[id] = quote

	return quote, nil
}

func (r *QuoteRepo) GetAll(ctx context.Context) ([]entity.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]entity.Quote, 0, len(r.quotes))

	for _, val := range r.quotes {
		result = append(result, val)
	}

	return result, nil
}

func (r *QuoteRepo) GetRand(ctx context.Context) (entity.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.quotes) == 0 {
		return entity.Quote{}, entity.ErrMissingQuotes
	}

	k := rand.Intn(len(r.quotes))

	for _, val := range r.quotes {
		if k == 0 {
			return val, nil
		}
		k--
	}

	return entity.Quote{}, fmt.Errorf("QuoteRepo - GetRand - rand.Intn")
}

func (r *QuoteRepo) GetByAuthor(ctx context.Context, author string) ([]entity.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]entity.Quote, 0)
	for _, q := range r.quotes {
		if q.Author == author {
			result = append(result, q)
		}
	}
	return result, nil
}

func (r *QuoteRepo) DeleteByID(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.quotes[id]; !exists {
		return entity.ErrNotFoundQuote
	}
	delete(r.quotes, id)
	return nil
}
