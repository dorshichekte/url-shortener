package grpcdbhandler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *DBHandler) Ping(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	err := h.useCase.Ping(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
