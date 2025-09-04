package grpcserver

import (
	"context"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	grpchandler "url-shortener/internal/app/adapter/primary/grpc/handler"
	"url-shortener/internal/app/adapter/primary/grpc/proto"
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/usecase"
	"url-shortener/internal/pkg/validator"
)

// New создание экземпляра GRPCAdapter
func New(logger *zap.Logger, config *config.Config, cases *usecase.UseCases, validator *validator.Validator) *Grpc {
	listener, err := net.Listen("tcp", config.Env.GrpcAddress)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	proto.RegisterUrlServer(server, grpchandler.New(logger, config.Env, cases, validator))

	startFunc := func() error {
		err = server.Serve(listener)
		return err
	}

	return &Grpc{
		start: startFunc,
	}
}

// Start старт grpc сервера
func (g *Grpc) Start(ctx context.Context) error {
	return g.start()
}
