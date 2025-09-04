package grpchandler

import (
	"context"
	grpcdbhandler "url-shortener/internal/app/adapter/primary/grpc/handler/db"
	grpcurlhandler "url-shortener/internal/app/adapter/primary/grpc/handler/url"
	"url-shortener/internal/app/adapter/primary/grpc/proto"
	config "url-shortener/internal/app/config/env"
	"url-shortener/internal/app/usecase"
	"url-shortener/internal/pkg/validator"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

// New конструктор grpc обработчиков
func New(logger *zap.Logger, config *config.Env, useCases *usecase.UseCases, validator *validator.Validator) *GRPCHandlers {
	return &GRPCHandlers{
		url: grpcurlhandler.New(useCases.URL, logger, validator, config),
		db:  grpcdbhandler.New(logger, useCases.DB),
	}
}

func (h *GRPCHandlers) AddBatch(ctx context.Context, req *proto.BatchRequest) (*proto.BatchResponse, error) {
	return h.url.AddBatch(ctx, req)
}

func (h *GRPCHandlers) AddShorten(ctx context.Context, req *proto.BatchRequest) (*proto.BatchResponse, error) {
	return h.url.AddShorten(ctx, req)
}

func (h *GRPCHandlers) DeleteBatch(ctx context.Context, req *proto.DeleteBatchRequest) (*emptypb.Empty, error) {
	return h.url.DeleteBatch(ctx, req)
}

func (h *GRPCHandlers) GetAllByUserID(ctx context.Context, empty *emptypb.Empty) (*proto.UrlResponse, error) {
	return h.url.GetAllByUserID(ctx, empty)
}

func (h *GRPCHandlers) GetStats(ctx context.Context, empty *emptypb.Empty) (*proto.ServiceStatsResponse, error) {
	return h.url.GetStats(ctx, empty)
}

func (h *GRPCHandlers) MakeFromJSON(ctx context.Context, req *proto.ShortenRequest) (*proto.ShortenResponse, error) {
	return h.url.MakeFromJSON(ctx, req)
}

func (h *GRPCHandlers) Ping(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	return h.db.Ping(ctx, empty)
}
