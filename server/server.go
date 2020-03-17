package server

import (
	"account-grpc"
	"context"
)

type accountRPCServer struct {
	account_grpc.UnimplementedAccountRoutesServer
}


// defined implementation of AccountRoutesServer interface
func (srv *accountRPCServer) SavePayment(ctx context.Context, payment *account_grpc.Payment) (*account_grpc.Reference, error){


}