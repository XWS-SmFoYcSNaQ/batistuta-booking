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

func (h RecommendationHandler) GetRecommendedAccommodations(ctx context.Context, request *recommendation.EmptyRequest) (*recommendation.RecommendedAccommodations_Response, error) {
	return h.RecommendationController.GetRecommendedAccommodations(ctx, request)
}
