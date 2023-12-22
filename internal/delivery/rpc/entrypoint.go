package rpc

import (
	searchPB "github.com/mar-coding/SearchEngineWrapper/APIs/proto-gen/services/search/v1"
	"github.com/mar-coding/SearchEngineWrapper/configs"
	"github.com/mar-coding/SearchEngineWrapper/internal/domain"
	"github.com/mar-coding/SearchEngineWrapper/pkg/errorHandler"
	"github.com/mar-coding/SearchEngineWrapper/pkg/logger"
	"google.golang.org/grpc"
)

type AllServices struct {
	searchPB.UnsafeSearchServiceServer

	useCases map[string]domain.Bridger
	logger   logger.Logger
	error    errorHandler.Handler
}

func NewSearchService(server *grpc.Server, errHandler errorHandler.Handler, logger logger.Logger, useCases ...domain.Bridger) {
	services := &AllServices{
		useCases: make(map[string]domain.Bridger),
		error:    errHandler,
		logger:   logger,
	}

	for _, useCase := range useCases {
		switch useCase.(type) {
		case domain.SearchUseCase:
			services.useCases[configs.SEARCH_KEY] = useCase
		}
	}

	searchPB.RegisterSearchServiceServer(server, services)
}
