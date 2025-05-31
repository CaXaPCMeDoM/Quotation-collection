package request

import "citatnik/internal/entity"

type Quote struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

func (q *Quote) ToEntity() entity.Quote {
	return entity.Quote{
		Author: q.Author,
		Text:   q.Quote,
	}
}
