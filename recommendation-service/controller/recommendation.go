package controller

import (
	"context"
	"recommendation-service/proto/recommendation"
	"recommendation-service/services"
)

type RecommendationController struct {
	RecommendationService *services.RecommendationService
}

func (c RecommendationController) Test(ctx context.Context, request *recommendation.Test_Request) (*recommendation.Test_Response, error) {
	return c.RecommendationService.Test()
}
