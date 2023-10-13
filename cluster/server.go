package cluster

import (
	"context"
	"fmt"
	"github.com/NoahAmethyst/simple-kube-operator/cluster/middleware"
	constant "github.com/NoahAmethyst/simple-kube-operator/constant"
	"github.com/NoahAmethyst/simple-kube-operator/opt_service"
	"github.com/NoahAmethyst/simple-kube-operator/protocol/pb/kube_opt_pb"
	"github.com/NoahAmethyst/simple-kube-operator/utils/log"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
	"runtime/debug"
	"strconv"
)

var (
	customFunc grpc_recovery.RecoveryHandlerFuncContext
)

var KubeOptServer *grpc.Server

func StartServer(grpcPort string) {
	if len(grpcPort) == 0 {
		grpcPort = strconv.Itoa(constant.DefaultGRPCPort)
	}

	grpcAddr := fmt.Sprintf("0.0.0.0:%s", grpcPort)
	lis, err := net.Listen("tcp", grpcAddr)

	if err != nil {
		log.Error().Fields(map[string]interface{}{"action": "grpc listener error", "error": err.Error()}).Send()
	}
	log.Info().Msgf("KubeOpt service start at address %s", grpcAddr)

	// Define customfunc to handle panic
	customFunc = func(ctx context.Context, p interface{}) error {
		log.Error().Msgf("[PANIC] %s\n\n%s", p, string(debug.Stack()))
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandlerContext(customFunc),
	}

	// Create a server. Recovery handlers should typically be last in the chain_info so that other middleware
	// (e.g. logging) can operate on the recovered state instead of being directly affected by any panic

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.LoggerInterceptor),
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(opts...),
			//otgrpc.OpenTracingServerInterceptor(thisTracer),

		),
		grpc.ChainStreamInterceptor(
			grpc_recovery.StreamServerInterceptor(opts...),
			//grpc_opentracing.StreamServerInterceptor(topts...),
		),
	)

	//register scaleSwap server
	kube_opt_pb.RegisterKubeOptServiceServer(grpcServer, opt_service.KubeOptServer{})
	KubeOptServer = grpcServer

	reflection.Register(grpcServer)

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Error().Err(err)
	}
}
