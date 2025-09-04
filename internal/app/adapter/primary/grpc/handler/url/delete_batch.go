package grpcurlhandler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	grpchandlerhelpers "url-shortener/internal/app/adapter/primary/grpc/handler/helpers"
	"url-shortener/internal/app/adapter/primary/grpc/proto"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	entity "url-shortener/internal/app/domain/entity/url"
	"url-shortener/internal/pkg/constants"
)

func (h *URLHandler) DeleteBatch(ctx context.Context, req *proto.DeleteBatchRequest) (*emptypb.Empty, error) {
	userID := grpchandlerhelpers.GetKeyFromContext(ctx, string(middleware.UserIDKey))
	if userID == "" {
		return nil, status.Error(codes.Internal, constants.ErrFailedGetUserIDFromContext.Error())
	}

	if len(req.ShortUrl) == 0 {
		return nil, status.Error(codes.InvalidArgument, "No urls provided to delete")
	}

	event := entity.DeleteBatch{
		ListURL: req.ShortUrl,
		UserID:  userID,
	}

	h.useCase.DeleteBatch(event)

	return &emptypb.Empty{}, nil
}
