package account_grpc

import (
	"account-grpc/internal/api"
	"account-grpc/internal/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	PORT = "2000"
)

func newServer(l *zap.SugaredLogger) *api.AccountRPCServer {
	accountGrpcServer := &api.AccountRPCServer{Logger: l}
	return accountGrpcServer
}

func main() {
	// logger initialization
	logger, err := config.BuildLogger()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	logger.Info("logger initialised")

	// created a new account rpc service
	accRPCService := newServer(logger)
	logger.Info("created account grpc server struct")

	// network listener creation
	logger.Info("network listener initialization")
	addr, err := net.LookupHost("grpc-account")
	if err != nil {
		logger.Error("host lookup failed:", err)
	}
	logger.Info("grpc server host", addr)
	lis, err := net.Listen("tcp", net.JoinHostPort(addr[0], PORT))
	if err != nil {
		logger.Errorw("failed to create network listener",
			"error", err)
	}
	logger.Info(lis.Addr().String(), " ", lis.Addr().Network())

	// grpc server initialization
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterAccountRoutesServer(grpcServer, accRPCService)
	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Errorw("Failed to create grpc server",
			"error", err)
	}
}
