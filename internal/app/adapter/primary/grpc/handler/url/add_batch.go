package grpcurlhandler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpchandlerhelpers "url-shortener/internal/app/adapter/primary/grpc/handler/helpers"
	"url-shortener/internal/app/adapter/primary/grpc/proto"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	entity "url-shortener/internal/app/domain/entity/url"
	"url-shortener/internal/pkg/constants"
)

func (h *URLHandler) AddBatch(ctx context.Context, req *proto.BatchRequest) (*proto.BatchResponse, error) {
	userID := grpchandlerhelpers.GetKeyFromContext(ctx, string(middleware.UserIDKey))
	if userID == "" {
		return nil, status.Error(codes.Internal, constants.ErrFailedGetUserIDFromContext.Error())
	}

	urls := make([]entity.Batch, len(req.Items))
	for i, item := range req.Items {
		urls[i] = entity.Batch{
			ShortURL:    item.CorrelationId,
			OriginalURL: item.OriginalUrl,
		}
	}

	result, err := h.useCase.AddBatch(ctx, urls, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var response proto.BatchResponse
	for _, item := range result {
		response.Items = append(response.Items, &proto.BatchResponse_Item{
			CorrelationId: item.ID,
			OriginalUrl:   item.OriginalURL,
		})
	}

	return &response, nil
}
