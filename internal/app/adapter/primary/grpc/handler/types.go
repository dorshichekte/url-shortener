package grpchandler

import (
	grpcdbhandler "url-shortener/internal/app/adapter/primary/grpc/handler/db"
	grpcurlhandler "url-shortener/internal/app/adapter/primary/grpc/handler/url"
	"url-shortener/internal/app/adapter/primary/grpc/proto"
)

// GRPCHandlers grpc обработчики
type GRPCHandlers struct {
	proto.UnimplementedUrlServer
	url *grpcurlhandler.URLHandler
	db  *grpcdbhandler.DBHandler
}
