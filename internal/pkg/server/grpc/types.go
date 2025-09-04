package grpcserver

// Grpc инкапсулирует GRPC-сервер и его зависимости.
type Grpc struct {
	start func() error
}
