package grpcurlhandler

import (
	"context"
	grpchandlerhelpers "url-shortener/internal/app/adapter/primary/grpc/handler/helpers"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	"url-shortener/internal/pkg/constants"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"url-shortener/internal/app/adapter/primary/grpc/proto"
)

func (h *URLHandler) AddShorten(ctx context.Context, req *proto.ShortenRequest) (*proto.ShortenResponse, error) {
	userID := grpchandlerhelpers.GetKeyFromContext(ctx, string(middleware.UserIDKey))
	if userID == "" {
		return nil, status.Error(codes.Internal, constants.ErrFailedGetUserIDFromContext.Error())
	}

	shortURL, err := h.useCase.AddShorten(ctx, req.Url, userID)
	baseURL := h.config.BaseURL
	fullURL := baseURL + "/" + shortURL
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var response = proto.ShortenResponse{
		Result: fullURL,
	}

	return &response, nil
}
