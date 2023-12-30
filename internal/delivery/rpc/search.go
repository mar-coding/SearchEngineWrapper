package rpc

import (
	"context"
	searchPB "github.com/mar-coding/SearchEngineWrapper/APIs/proto-gen/services/search/v1"
	"github.com/mar-coding/SearchEngineWrapper/configs"
	"github.com/mar-coding/SearchEngineWrapper/internal/domain"
)

func (a AllServices) ListItemsSearch(ctx context.Context, request *searchPB.ListItemsSearchRequest) (*searchPB.ListItemsSearchResponse, error) {
	useCase := domain.Bridge[domain.SearchUseCase](configs.SEARCH_KEY, a.useCases)

	items, err := useCase.ListItems(ctx, request.Q, request.PageSize, request.PageNo, false)
	if err != nil {
		return nil, err
	}

	itemsPB := make([]*searchPB.Item, 0)
	for _, item := range items.Items {
		temp := new(searchPB.Item)
		(*temp).Url = item.URL
		(*temp).Title = item.Title
		(*temp).Description = item.Description
		itemsPB = append(itemsPB, temp)
	}

	return &searchPB.ListItemsSearchResponse{
		Items:           itemsPB,
		PageSize:        items.PageSize,
		PageNo:          items.PageNo,
		TotalItemsCount: items.TotalItemsCount,
	}, nil
}
