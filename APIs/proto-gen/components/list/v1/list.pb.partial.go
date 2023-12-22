package list

import (
	"github.com/mar-coding/SearchEngineWrapper/pkg/helper"
	"github.com/mar-coding/SearchEngineWrapper/pkg/mongodb/pagination"
)

type PaginatedList pagination.List

// NewResponse create paginated response for convert to proto
func NewResponse(paginatedList *pagination.List) *PaginatedList {
	return (*PaginatedList)(paginatedList)
}

// AsModel convert PaginatedListRequest to List
func (x *PaginatedListRequest) AsModel() *pagination.List {
	l, _ := helper.ConvertProtoToModel[*pagination.List](x)

	if l.PageSize == 0 {
		l.PageSize = 10
	}

	if l.PageNo == 0 {
		l.PageNo = 1
	}

	return l
}

// AsProto convert PaginatedList to PaginatedListResponse
func (p *PaginatedList) AsProto() *PaginatedListResponse {
	l, _ := helper.ConvertModelToProto[*PaginatedListResponse](p)
	return l
}
