package handlers

import (
	"context"
	"recommendation-service/controller"
	"recommendation-service/proto/recommendation"
)

type RecommendationHandler struct {
	recommendation.UnimplementedRecommendationServiceServer
	RecommendationController *controller.RecommendationController
}

func (h RecommendationHandler) Test(ctx context.Context, request *recommendation.Test_Request) (*recommendation.Test_Response, error) {
	return h.RecommendationController.Test(ctx, request)
}
