package grpcadapter

import server "url-shortener/internal/pkg/server/grpc"

// GRPCAdapter сервер grpcadapter
type GRPCAdapter struct {
	server *server.Grpc
}
