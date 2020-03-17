package main

import (
	"account-grpc/server"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net"
)

const (
	PORT = 3000
)

func newServer(l *zap.SugaredLogger) *server.AccountRPCServer {
	accountGrpcServer := &server.AccountRPCServer{Logger:l}
	return accountGrpcServer
}

func main() {
	// logger init
	logger, err := server.BuildLogger()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	logger.Info("logger initialised")
	acgs := newServer(logger)
	logger.Info("created account grpc server struct")

	logger.Info("network listener init")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", PORT))
	if err != nil {
		logger.Errorw("failed to create network listener",
			"error", err)
	}

}
