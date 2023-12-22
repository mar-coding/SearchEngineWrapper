package usecase

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/mar-coding/SearchEngineWrapper/configs"
	"github.com/mar-coding/SearchEngineWrapper/internal/domain"
	"github.com/mar-coding/SearchEngineWrapper/pkg/elastic"
	"github.com/mar-coding/SearchEngineWrapper/pkg/errorHandler"
	"github.com/mar-coding/SearchEngineWrapper/pkg/logger"
)

type SearchUseCase struct {
	logger  logger.Logger
	error   errorHandler.Handler
	elastic *elastic.Client
}

func NewSearch(elastic *elastic.Client, logger logger.Logger, error errorHandler.Handler) domain.SearchUseCase {
	return &SearchUseCase{
		error:   error,
		logger:  logger,
		elastic: elastic,
	}
}

func (s *SearchUseCase) None() {}

func (s *SearchUseCase) ListItems(ctx context.Context, query string, pageSize, pageNo int32, withWildCard bool) (domain.SearchResult, error) {
	var searchResult domain.SearchResult
	results := make([]*domain.Item, 0)
	resultCount := 0

	query = strings.TrimSpace(query)
	if withWildCard {
		query = fmt.Sprintf("*%s*", query)
	}

	if pageSize == 0 {
		pageSize = configs.ELASTIC_DEFAULT_PAGE_SIZE
	}

	if pageNo == 0 {
		pageNo = configs.ELASTIC_DEFAULT_PAGE_NUMBER
	}
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		//count, err := s.elastic.ExecuteQuery(ctx, configs.ELASTIC_ITEM_INDEX_KEY, []string{
		//	"title",
		//	"description",
		//	"url",
		//}, query, pageNo, pageSize, &results)
		count, err := mockExecuteQuery(ctx, configs.ELASTIC_ITEM_INDEX_KEY, []string{
			"title",
			"description",
			"url",
		}, query, pageNo, pageSize, &results)
		if err != nil {
			s.logger.ErrorContext(ctx, true, err.Error())
		}
		resultCount += count
	}()

	wg.Wait()

	searchResult.Items = results
	searchResult.PageSize = pageSize
	searchResult.PageNo = pageNo
	searchResult.TotalItemsCount = int32(resultCount)

	return searchResult, nil
}

func mockExecuteQuery(_ context.Context, _ string, _ []string, _ string, _ int32, _ int32, response any) (int, error) {
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

	items := make([]*domain.Item, 0)
	items = append(items, item1)
	items = append(items, item2)

	if res, ok := response.(*[]*domain.Item); ok {
		*res = items
	}

	return 2, nil
}
