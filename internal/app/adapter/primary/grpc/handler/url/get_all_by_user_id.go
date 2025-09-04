package grpcurlhandler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	grpchandlerhelpers "url-shortener/internal/app/adapter/primary/grpc/handler/helpers"
	"url-shortener/internal/app/adapter/primary/grpc/proto"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	"url-shortener/internal/pkg/constants"
)

func (h *URLHandler) GetAllByUserID(ctx context.Context, _ *emptypb.Empty) (*proto.UserUrlResponse, error) {
	userID := grpchandlerhelpers.GetKeyFromContext(ctx, string(middleware.UserIDKey))
	if userID == "" {
		return nil, status.Error(codes.Internal, constants.ErrFailedGetUserIDFromContext.Error())
	}

	urls, err := h.useCase.GetAllByUserID(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if len(urls) == 0 {
		return &proto.UserUrlResponse{}, nil
	}

	var response proto.UserUrlResponse
	for _, url := range urls {
		response.Urls = append(response.Urls, &proto.UrlResponse{
			ShortUrl:    url.ShortURL,
			OriginalUrl: url.OriginalURL,
		})
	}

	return &response, nil
}
