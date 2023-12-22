package domain

import (
	"context"
)

type SearchUseCase interface {
	Bridger
	ListItems(ctx context.Context, query string, pageSize, pageNo int32, withWildCard bool) (SearchResult, error)
}
