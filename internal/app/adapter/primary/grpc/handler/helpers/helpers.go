package grpchandlerhelpers

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func GetKeyFromContext(ctx context.Context, key string) string {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	ids := meta.Get(key)
	if len(ids) == 0 {
		return ""
	}

	return ids[0]
}
