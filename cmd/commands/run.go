package commands

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mar-coding/SearchEngineWrapper/APIs"
	searchPB "github.com/mar-coding/SearchEngineWrapper/APIs/proto-gen/services/search/v1"
	"github.com/mar-coding/SearchEngineWrapper/configs"
	"github.com/mar-coding/SearchEngineWrapper/internal/app/search"
	"github.com/mar-coding/SearchEngineWrapper/pkg/elastic"
	"github.com/mar-coding/SearchEngineWrapper/pkg/middlewares"
	"github.com/mar-coding/SearchEngineWrapper/pkg/transport"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

const (
	headerSignatureQuery = "X-Signature-Query"
	headerCaptchaQuery   = "X-Captcha-Key"
	headerServiceName    = "X-Service-Name"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run Application",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := configs.NewConfig(configPath)
		if err != nil {
			return err
		}

		info, err := configs.NewServiceInfo()
		if err != nil {
			return err
		}

		errHandler, err := configs.NewError(info, cfg)
		if err != nil {
			return err
		}

		logging, err := configs.NewLogger(cfg, info)
		if err != nil {
			return err
		}

		grpcAddr := fmt.Sprintf("%s:%s", cfg.Address, strconv.Itoa(cfg.Grpc.Port))
		grpcServer, err := transport.NewGRPCServer(
			grpcAddr,
			cfg.Development,
			middlewares.GrpcRecovery(logging),
			middlewares.GrpcSentryPerformance(logging.GetSentryClient()),
			middlewares.GrpcValidator(errHandler),
			middlewares.GrpcJwtMiddleware(permissionOptions, info, errHandler, os.Getenv("PUBLIC_SECRET"),
				os.Getenv("PRIVATE_SECRET"), ""),
			middlewares.GRPCLogging(logging),
		)
		if err != nil {
			return err
		}

		httpAddr := fmt.Sprintf("%s:%s", cfg.Address, strconv.Itoa(cfg.Rest.Port))
		httpServer := transport.NewHTTPServer(
			cmd.Context(),
			httpAddr,
			grpcAddr,
			cfg.Development,
			APIs.Swagger,
			[]string{headerSignatureQuery, headerCaptchaQuery, headerServiceName, "Content-Type", "Content-Disposition"},
			cfg.Origins,
			defaultHttpMiddlewareFunc,
			runtime.WithIncomingHeaderMatcher(headers),
			runtime.WithErrorHandler(middlewares.ErrorHandler),
		)

		elasticsearch, err := elastic.NewElasticSearch(cmd.Context(), cfg.Database.Elastic.Addresses, cfg.Database.Elastic.Username, cfg.Database.Elastic.Password, cfg.Development)
		if err != nil {
			return err
		} else {
			logging.InfoContext(context.Background(), false, "connected to ElasticSearch")
		}

		if err := httpServer.RegisterServiceEndpoint(searchPB.RegisterSearchServiceHandlerFromEndpoint); err != nil {
			return err
		}

		application, err := search.NewApp(cmd.Context(),
			grpcServer, httpServer,
			cfg,
			info,
			errHandler,
			logging,
			elasticsearch,
		)
		if err != nil {
			return err
		}

		application.Run(cmd.Context())
		return nil
	},
}

func headers(key string) (string, bool) {
	switch key {
	case headerCaptchaQuery:
		return key, true
	case headerServiceName:
		return key, true
	case headerSignatureQuery:
		return key, true
	default:
		return key, false
	}
}

func permissionOptions(methodFullName string) ([]int32, bool, bool, bool, error) {
	desc, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(methodFullName))
	if err != nil {
		return nil, false, false, false, err
	}

	methodDesc, ok := desc.(protoreflect.MethodDescriptor)
	if !ok {
		return nil, false, false, false, errors.New("failed to assert MethodDescriptor")
	}

	options, ok := methodDesc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return nil, false, false, false, errors.New("failed to assert MethodOptions")
	}

	permExt := proto.GetExtension(options, searchPB.E_Permission)

	perm, ok := permExt.(*searchPB.Permission)
	if !ok {
		return nil, false, false, false, errors.New("failed to assertion captcha object with proto extension")
	}

	permsCode := make([]int32, 0, len(perm.Permissions))
	for _, permission := range perm.Permissions {
		permsCode = append(permsCode, int32(permission))
	}

	return permsCode, perm.Optional, perm.ValidatePermissions, perm.Captcha, nil
}

func defaultHttpMiddlewareFunc(handler http.Handler) http.Handler {
	return handler
}
