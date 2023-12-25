package mocking

import (
	"context"
	"github.com/mar-coding/SearchEngineWrapper/internal/domain"
)

func MockExecuteQuery(_ context.Context, _ string, _ []string, _ string, _ int32, _ int32, response any) (int, error) {
	item1 := &domain.Item{
		Title:       "Item 1",
		Description: "Description 1",
		Url:         "https://example.com/item1",
	}

	item2 := &domain.Item{
		Title:       "Item 2",
		Description: "Description 2",
		Url:         "https://example.com/item2",
	}

	item3 := &domain.Item{
		Title:       "Item 3",
		Description: "Description 3",
		Url:         "https://example.com/item3",
	}

	item4 := &domain.Item{
		Title:       "Item 4",
		Description: "Description 4",
		Url:         "https://example.com/item4",
	}

	item5 := &domain.Item{
		Title:       "Item 5",
		Description: "Description 5",
		Url:         "https://example.com/item5",
	}

	item6 := &domain.Item{
		Title:       "Item 6",
		Description: "Description 6",
		Url:         "https://example.com/item6",
	}

	item7 := &domain.Item{
		Title:       "Item 7",
		Description: "Description 7",
		Url:         "https://example.com/item7",
	}

	item8 := &domain.Item{
		Title:       "Item 8",
		Description: "Description 8",
		Url:         "https://example.com/item8",
	}

	items := make([]*domain.Item, 0)
	items = append(items, item1)
	items = append(items, item2)
	items = append(items, item3)
	items = append(items, item4)
	items = append(items, item5)
	items = append(items, item6)
	items = append(items, item7)
	items = append(items, item8)

	if res, ok := response.(*[]*domain.Item); ok {
		*res = items
	}

	return len(items), nil
}
