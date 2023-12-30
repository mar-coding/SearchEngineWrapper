package usecase

import (
	"context"
	"fmt"
	"github.com/mar-coding/SearchEngineWrapper/configs"
	"github.com/mar-coding/SearchEngineWrapper/internal/domain"
	"github.com/mar-coding/SearchEngineWrapper/pkg/elastic"
	"github.com/mar-coding/SearchEngineWrapper/pkg/errorHandler"
	"github.com/mar-coding/SearchEngineWrapper/pkg/logger"
	"strings"
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
	var elasticResult []domain.ElasticResult
	var result domain.SearchResult
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
	count, err := s.elastic.ExecuteQuery(ctx, configs.ELASTIC_ITEM_INDEX_KEY, []string{
		"title",
		"text",
		"url",
	}, query, pageNo, pageSize, &elasticResult)
	//count, err := mocking.MockExecuteQuery(ctx, configs.ELASTIC_ITEM_INDEX_KEY, []string{
	//	"title",
	//	"description",
	//	"url",
	//}, query, pageNo, pageSize, &results)
	if err != nil {
		s.logger.ErrorContext(ctx, true, err.Error())
	}
	resultCount += count

	itemArr := make([]*domain.Item, 0)

	for i := 0; i < len(elasticResult); i++ {
		temp := elasticResult[i].Source
		temp.Description = lenLimiter(cleanString(strings.Join(temp.Text, " ")), 200)
		itemArr = append(itemArr, &temp)
	}

	result.Items = itemArr
	result.PageSize = pageSize
	result.PageNo = pageNo
	result.TotalItemsCount = int32(resultCount)

	return result, nil
}

func cleanString(input string) string {
	cleaned := strings.ReplaceAll(input, "\n", "")
	cleaned = strings.ReplaceAll(cleaned, "\t", "")
	cleaned = strings.ReplaceAll(cleaned, "\r", "")
	return cleaned
}

func lenLimiter(input string, limit int) string {
	result := []rune(input)

	if len(result) > limit {
		output := string(result[:limit]) + " ..."
		output = strings.TrimSpace(output)
		return output
	}

	output := strings.TrimSpace(string(result))
	return output
}
