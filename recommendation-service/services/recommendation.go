package services

import (
	"context"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/accommodation"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/user"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"recommendation-service/infrastructure/database"
)

type RecommendationService struct {
	RecommendationRepo *database.RecommendationRepo
}

func (s RecommendationService) SetDataForDb(accommodations []*accommodation.AccommodationDTO, users []*user.User) {
	ctx := context.Background()
	session := s.RecommendationRepo.Driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)
	// Insert accommodations
	err := s.insertAccommodations(session, accommodations, ctx)
	if err != nil {
		s.RecommendationRepo.Logger.Println("Error inserting accommodations:", err)
		// Handle the error accordingly
		return
	}

	// Insert users
	err = s.insertUsers(session, users, ctx)
	if err != nil {
		s.RecommendationRepo.Logger.Println("Error inserting users:", err)
		// Handle the error accordingly
		return
	}

	// Insert ratings
	err = s.insertRatings(session, accommodations, users, ctx)
	if err != nil {
		s.RecommendationRepo.Logger.Println("Error inserting ratings:", err)
		// Handle the error accordingly
		return
	}
}

func (s RecommendationService) insertAccommodations(session neo4j.SessionWithContext, accommodations []*accommodation.AccommodationDTO, ctx context.Context) error {
	_, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (interface{}, error) {
			for _, acc := range accommodations {
				result, err := transaction.Run(ctx,
					"CREATE (a:Accommodation) SET a.id = $id, a.name = $name, a.benefits = $benefits, a.minGuests = $minGuests, a.maxGuests = $maxGuests, a.basePrice = $basePrice",
					map[string]interface{}{
						"id":        acc.Id,
						"name":      acc.Name,
						"benefits":  acc.Benefits,
						"minGuests": acc.MinGuests,
						"maxGuests": acc.MaxGuests,
						"basePrice": acc.BasePrice,
					})
				if err != nil {
					return nil, err
				}
				if result.Err() != nil {
					return nil, result.Err()
				}
			}

			return nil, nil
		})

	return err
}

func (s RecommendationService) insertUsers(session neo4j.SessionWithContext, users []*user.User, ctx context.Context) error {
	_, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (interface{}, error) {
			for _, usr := range users {
				if usr.Role == user.UserRole_Guest {
					result, err := transaction.Run(ctx,
						"CREATE (u:User) SET u.id = $id, u.role = $role, u.username = $username, u.firstName = $firstName, u.lastName = $lastName, u.email = $email, u.livingPlace = $livingPlace",
						map[string]interface{}{
							"id":          usr.Id,
							"role":        usr.Role,
							"username":    usr.Username,
							"firstName":   usr.FirstName,
							"lastName":    usr.LastName,
							"email":       usr.Email,
							"livingPlace": usr.LivingPlace,
						})
					if err != nil {
						return nil, err
					}
					if result.Err() != nil {
						return nil, result.Err()
					}
				}
			}

			return nil, nil
		})

	return err
}

func (s RecommendationService) insertRatings(session neo4j.SessionWithContext, accommodations []*accommodation.AccommodationDTO, users []*user.User, ctx context.Context) error {
	_, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (interface{}, error) {
			for _, acc := range accommodations {
				for _, rating := range acc.Ratings {
					result, err := transaction.Run(ctx,
						"CREATE (r:Rating) SET r.id = $ratingId, r.value = $ratingValue",
						map[string]interface{}{
							"ratingId":    rating.Id,
							"ratingValue": rating.Value,
						})
					if err != nil {
						return nil, err
					}
					if result.Err() != nil {
						return nil, result.Err()
					}

					// Link Rating node with Accommodation node
					_, err = transaction.Run(ctx,
						"MATCH (a:Accommodation), (r:Rating) WHERE id(a) = $accommodationId AND r.id = $ratingId CREATE (a)-[:HAS_RATING]->(r)",
						map[string]interface{}{
							"accommodationId": acc.Id,
							"ratingId":        rating.Id,
						})
					if err != nil {
						return nil, err
					}

					// Link Rating node with User node
					for _, user := range users {
						if user.Id == rating.UserId {
							_, err = transaction.Run(ctx,
								"MATCH (u:User), (r:Rating) WHERE u.id = $userId AND r.id = $ratingId CREATE (u)-[:HAS_RATING]->(r)",
								map[string]interface{}{
									"userId":   user.Id,
									"ratingId": rating.Id,
								})
							if err != nil {
								return nil, err
							}
							break
						}
					}

				}
			}

			return nil, nil
		})

	return err
}
