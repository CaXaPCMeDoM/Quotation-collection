package usecase

import (
	"citatnik/internal/entity"
	"context"
)

type (
	Quote interface {
		Add(ctx context.Context, quote entity.Quote) (entity.Quote, error)
		GetAll(ctx context.Context) ([]entity.Quote, error)
		GetByAuthor(ctx context.Context, author string) ([]entity.Quote, error)
		GetRandom(ctx context.Context) (entity.Quote, error)
		DeleteByID(ctx context.Context, id string) error
	}
)
