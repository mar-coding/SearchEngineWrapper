package usecase

import (
	"context"
	"github.com/mar-coding/SearchEngineWrapper/internal/domain/mocking"
	"testing"

	"github.com/mar-coding/SearchEngineWrapper/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestMockExecuteQuery(t *testing.T) {
	res := make([]*domain.Item, 0)
	searchResult, err := mocking.MockExecuteQuery(context.Background(), "nil", nil, "nil", 0, 0, &res)
	if err != nil {
		assert.Errorf(t, err, "got error")
	}
	assert.Equal(t, 2, searchResult)

	// Assert the items in res
	assert.Equal(t, 2, len(res))
	assert.Equal(t, "Item 1", res[0].Title)
	assert.Equal(t, "Description 1", res[0].Description)
	assert.Equal(t, "https://example.com/item1", res[0].Url)
	assert.Equal(t, "Item 2", res[1].Title)
	assert.Equal(t, "Description 2", res[1].Description)
	assert.Equal(t, "https://example.com/item2", res[1].Url)
}
