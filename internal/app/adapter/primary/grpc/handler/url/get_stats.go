package grpcurlhandler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"url-shortener/internal/app/adapter/primary/grpc/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *URLHandler) GetStats(ctx context.Context, _ *emptypb.Empty) (*proto.ServiceStatsResponse, error) {
	stats, err := h.useCase.GetStats(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &proto.ServiceStatsResponse{
		UserCount: int64(stats.UserCount),
		Urls:      int64(stats.URLCount),
	}

	return resp, nil
}
