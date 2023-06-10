package common

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
)

func GetAuthHeader(ctx *context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(*ctx)
	if !ok {
		return "", errors.New("failed to read metadata")
	}

	authHeaders := md.Get("Authorization")
	if len(authHeaders) == 0 {
		return "", errors.New("missing authorization header")
	}

	return authHeaders[0], nil
}
