package repo

import (
	"citatnik/internal/entity"
	"context"
)

type (
	QuoteRepo interface {
		Create(ctx context.Context, quote entity.Quote) (entity.Quote, error)
		GetAll(ctx context.Context) ([]entity.Quote, error)
		GetRand(ctx context.Context) (entity.Quote, error)
		GetByAuthor(ctx context.Context, author string) ([]entity.Quote, error)
		DeleteByID(ctx context.Context, id string) error
	}
)
