package search

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mar-coding/SearchEngineWrapper/configs"
	"github.com/mar-coding/SearchEngineWrapper/internal/delivery/rpc"
	"github.com/mar-coding/SearchEngineWrapper/internal/usecase"
	"github.com/mar-coding/SearchEngineWrapper/pkg/configHandler"
	"github.com/mar-coding/SearchEngineWrapper/pkg/elastic"
	"github.com/mar-coding/SearchEngineWrapper/pkg/errorHandler"
	"github.com/mar-coding/SearchEngineWrapper/pkg/logger"
	"github.com/mar-coding/SearchEngineWrapper/pkg/serviceInfo"
	"github.com/mar-coding/SearchEngineWrapper/pkg/transport"
)

type AppBootstrapper interface {
	Run(ctx context.Context)
	Shutdown(ctx context.Context) error
	GetServiceInfo() *serviceInfo.ServiceInfo
	GetServiceConfig() *configHandler.BaseConfig[configs.ExtraData]
	GetErrorHandler() errorHandler.Handler
}

type Application struct {
	grpcServer transport.GRPCBootstrapper
	httpServer transport.HTTPBootstrapper

	serviceInfo   *serviceInfo.ServiceInfo
	serviceConfig *configHandler.BaseConfig[configs.ExtraData]
	logger        logger.Logger
	error         errorHandler.Handler
	elastic       *elastic.Client
}

func NewApp(
	_ context.Context,
	grpcServer transport.GRPCBootstrapper,
	httpServer transport.HTTPBootstrapper,
	serviceConfig *configHandler.BaseConfig[configs.ExtraData],
	serviceInfo *serviceInfo.ServiceInfo,
	errorHandler errorHandler.Handler,
	logger logger.Logger,
	elastic *elastic.Client,
) (AppBootstrapper, error) {
	app := new(Application)

	if serviceConfig == nil {
		return nil, errors.New("service config is nil")
	}

	app.serviceConfig = serviceConfig
	app.serviceInfo = serviceInfo
	app.logger = logger
	app.error = errorHandler
	app.grpcServer = grpcServer
	app.httpServer = httpServer
	app.elastic = elastic

	return app, nil
}

func (a *Application) Run(ctx context.Context) {
	if err := a.registerServiceLayers(ctx); err != nil {
		log.Fatal(err)
	}

	a.grpcServer.Start()
	a.logger.InfoContext(ctx, false, "grpc server on has been started", "address",
		fmt.Sprintf("%s:%v", a.serviceConfig.Address, a.serviceConfig.Grpc.Port))

	a.httpServer.Start()
	a.logger.InfoContext(ctx, false, "http server on has been started",
		"address", fmt.Sprintf("%s:%v", a.serviceConfig.Address, a.serviceConfig.Rest.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		if err := a.Shutdown(ctx); err != nil {
			panic(err)
		}
		a.logger.WarnContext(ctx, false, "app/run: signal received", "signal", s.String())
	case err := <-a.httpServer.Notify():
		if err := a.Shutdown(ctx); err != nil {
			panic(err)
		}
		a.logger.WarnContext(ctx, false, "app/run/httpServer.Notify: ", "error", err)
	case err := <-a.grpcServer.Notify():
		if err := a.Shutdown(ctx); err != nil {
			panic(err)
		}
		a.logger.WarnContext(ctx, false, "app/run/grpcServer.Notify: ", "error", err)
	}
}

func (a *Application) Shutdown(ctx context.Context) error {
	if err := a.grpcServer.Shutdown(ctx); err != nil {
		return err
	}

	if err := a.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func (a *Application) GetServiceInfo() *serviceInfo.ServiceInfo {
	return a.serviceInfo
}

func (a *Application) GetServiceConfig() *configHandler.BaseConfig[configs.ExtraData] {
	return a.serviceConfig
}

func (a *Application) GetErrorHandler() errorHandler.Handler {
	return a.error
}

func (a *Application) registerServiceLayers(_ context.Context) error {
	searchUseCase := usecase.NewSearch(a.elastic, a.logger, a.error)

	rpc.NewSearchService(a.grpcServer.GetGRPCServer(), a.error, a.logger,
		searchUseCase,
	)
	return nil
}
