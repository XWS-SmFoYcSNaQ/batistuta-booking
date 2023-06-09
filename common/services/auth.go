package services

import (
	"context"
	"errors"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/auth"
	"google.golang.org/grpc/metadata"
)

type AuthService struct {
	AuthClient *auth.AuthServiceClient
}

func NewAuthService(authClient *auth.AuthServiceClient) *AuthService {
	return &AuthService{
		AuthClient: authClient,
	}
}

func (s AuthService) ValidateToken(ctx *context.Context) (*auth.Verify_Response, error) {
	md, ok := metadata.FromIncomingContext(*ctx)
	if !ok {
		return nil, errors.New("failed to read metadata")
	}

	authHeaders := md.Get("Authorization")
	if len(authHeaders) == 0 {
		return nil, errors.New("missing authorization header")
	}

	authHeader := authHeaders[0]
	authCtx := metadata.AppendToOutgoingContext(*ctx, "Authorization", authHeader)
	res, err := (*s.AuthClient).Verify(authCtx, &auth.Empty_Request{})
	if err != nil || !res.Verified {
		return nil, errors.New("missing authorization header")
	}

	return res, err
}
