package rpc

import (
	"context"
	searchPB "github.com/mar-coding/SearchEngineWrapper/APIs/proto-gen/services/search/v1"
	"github.com/mar-coding/SearchEngineWrapper/configs"
	"github.com/mar-coding/SearchEngineWrapper/internal/domain"
	"github.com/mar-coding/SearchEngineWrapper/pkg/helper"
)

func (a AllServices) ListItemsSearch(ctx context.Context, request *searchPB.ListItemsSearchRequest) (*searchPB.ListItemsSearchResponse, error) {
	useCase := domain.Bridge[domain.SearchUseCase](configs.SEARCH_KEY, a.useCases)

	items, err := useCase.ListItems(ctx, "", 10, 1, false)
	if err != nil {
		return nil, err
	}

	itemsPB := make([]*searchPB.Item, 0)
	for _, item := range items.Items {
		itemPB, err := helper.ConvertModelToProto[*searchPB.Item](item)
		if err != nil {
			return nil, err
		}
		if itemPB.Id != "0" {
			itemsPB = append(itemsPB, itemPB)
		}
	}

	return &searchPB.ListItemsSearchResponse{
		Items:           itemsPB,
		PageSize:        items.PageSize,
		PageNo:          items.PageNo,
		TotalItemsCount: items.TotalItemsCount,
	}, nil

}
