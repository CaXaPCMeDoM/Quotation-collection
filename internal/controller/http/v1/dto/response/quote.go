package response

import "citatnik/internal/entity"

type Quote struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

func (q *Quote) ToResp(quote entity.Quote) {
	q.Author = quote.Author
	q.Quote = quote.Text
}
