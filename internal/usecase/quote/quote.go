package quote

import (
	"citatnik/internal/entity"
	"citatnik/internal/repo"
	"context"
)

type UseCase struct {
	quoteRepo repo.QuoteRepo
}

func New(quoteRepo repo.QuoteRepo) *UseCase {
	return &UseCase{
		quoteRepo: quoteRepo,
	}
}

func (uc *UseCase) Add(ctx context.Context, quote entity.Quote) (entity.Quote, error) {
	return uc.quoteRepo.Create(ctx, quote)
}

func (uc *UseCase) GetAll(ctx context.Context) ([]entity.Quote, error) {
	return uc.quoteRepo.GetAll(ctx)
}

func (uc *UseCase) GetByAuthor(ctx context.Context, author string) ([]entity.Quote, error) {
	return uc.quoteRepo.GetByAuthor(ctx, author)
}

func (uc *UseCase) GetRandom(ctx context.Context) (entity.Quote, error) {
	return uc.quoteRepo.GetRand(ctx)
}

func (uc *UseCase) DeleteByID(ctx context.Context, id string) error {
	return uc.quoteRepo.DeleteByID(ctx, id)
}
