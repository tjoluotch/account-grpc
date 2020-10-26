package api

import (
	"context"
	"go.uber.org/zap"
	"sync"
)

type AccountRPCServer struct {
	UnimplementedAccountRoutesServer
	mu sync.Mutex
	Logger *zap.SugaredLogger
}


// defined implementation of AccountRoutesServer interface
func (srv *AccountRPCServer) SavePayment(ctx context.Context, payment *Payment) (*Reference, error){
	srv.Logger.Info("Account grpc server called stub")
	return &Reference{}, nil
}