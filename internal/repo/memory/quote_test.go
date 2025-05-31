package memory

import (
	"citatnik/internal/entity"
	"context"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestQuoteRepo_CreateAndGetAll(t *testing.T) {
	repo := New()
	ctx := context.Background()

	q1, err := repo.Create(ctx, entity.Quote{Author: "A", Text: "Q1"})
	assert.NoError(t, err)
	q2, err := repo.Create(ctx, entity.Quote{Author: "B", Text: "Q2"})
	assert.NoError(t, err)
	q3, err := repo.Create(ctx, entity.Quote{Author: "A", Text: "Q3"})
	assert.NoError(t, err)

	assert.NotEmpty(t, q1.ID)
	assert.NotEmpty(t, q2.ID)
	assert.NotEmpty(t, q3.ID)
	assert.NotEqual(t, q1.ID, q2.ID)
	assert.NotEqual(t, q2.ID, q3.ID)

	all, err := repo.GetAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, all, 3)

	ids := map[string]bool{q1.ID: true, q2.ID: true, q3.ID: true}
	for _, q := range all {
		assert.True(t, ids[q.ID], "unexpected quote ID %s", q.ID)
	}
}

func TestQuoteRepo_GetByAuthor(t *testing.T) {
	repo := New()
	ctx := context.Background()

	_, _ = repo.Create(ctx, entity.Quote{Author: "X", Text: "Q1"})
	_, _ = repo.Create(ctx, entity.Quote{Author: "Y", Text: "Q2"})
	q3, _ := repo.Create(ctx, entity.Quote{Author: "X", Text: "Q3"})

	res, err := repo.GetByAuthor(ctx, "X")
	assert.NoError(t, err)
	assert.Len(t, res, 2)

	for _, q := range res {
		assert.Equal(t, "X", q.Author)
	}

	found := false
	for _, q := range res {
		if q.ID == q3.ID {
			found = true
		}
	}
	assert.True(t, found, "expected to find quote ID %s", q3.ID)
}

func TestQuoteRepo_DeleteByID(t *testing.T) {
	repo := New()
	ctx := context.Background()

	q, _ := repo.Create(ctx, entity.Quote{Author: "D", Text: "ToDelete"})

	err := repo.DeleteByID(ctx, q.ID)
	assert.NoError(t, err)

	all, _ := repo.GetAll(ctx)
	for _, v := range all {
		assert.NotEqual(t, q.ID, v.ID)
	}

	err = repo.DeleteByID(ctx, "nonexistent")
	assert.Equal(t, entity.ErrNotFoundQuote, err)
}

func TestQuoteRepo_GetRand_Empty(t *testing.T) {
	repo := New()
	ctx := context.Background()

	_, err := repo.GetRand(ctx)
	assert.Equal(t, entity.ErrMissingQuotes, err)
}

func TestQuoteRepo_GetRand_NonEmpty(t *testing.T) {
	repo := New()
	ctx := context.Background()

	q, _ := repo.Create(ctx, entity.Quote{Author: "R", Text: "OnlyOne"})

	rand.Seed(42)
	got, err := repo.GetRand(ctx)
	assert.NoError(t, err)
	assert.Equal(t, q.ID, got.ID)

	q2, _ := repo.Create(ctx, entity.Quote{Author: "R", Text: "Second"})
	rand.Seed(time.Now().UnixNano())
	got2, err := repo.GetRand(ctx)
	assert.NoError(t, err)
	assert.Contains(t, []string{q.ID, q2.ID}, got2.ID)
}
