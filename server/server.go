package server

import (
	"account-grpc"
	"context"
	"go.uber.org/zap"
	"sync"
)

type AccountRPCServer struct {
	account_grpc.UnimplementedAccountRoutesServer
	mu sync.Mutex
	Logger *zap.SugaredLogger
}


// defined implementation of AccountRoutesServer interface
func (srv *AccountRPCServer) SavePayment(ctx context.Context, payment *account_grpc.Payment) (*account_grpc.Reference, error){
	srv.Logger.Info("Account grpc server called stub")
	return &account_grpc.Reference{}, nil
}