package services

import (
	"context"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/accommodation"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/booking"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/rating"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/user"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"log"
	"recommendation-service/infrastructure/database"
	"recommendation-service/proto/recommendation"
	"time"
)

type RecommendationService struct {
	RecommendationRepo *database.RecommendationRepo
}

func (s RecommendationService) GetRecommendedAccommodations(userID string) ([]*recommendation.RecommendedAccommodationDTO, error) {
	ctx := context.Background()
	session := s.RecommendationRepo.Driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	// Cypher query to get recommended accommodations for the given user ID
	query := `
		MATCH (u:User)-[:HAS_RESERVATION]->(:Reservation)<-[:HAS_RESERVATION]-(a:Accommodation)
		WHERE u.id = $userID
		WITH COLLECT(DISTINCT a) AS userAccommodations
		
		MATCH (a:Accommodation)
		WHERE NOT a IN userAccommodations
		WITH userAccommodations, COLLECT(a) AS otherAccommodations
		
		MATCH (otherUser:User)-[:HAS_RESERVATION]->(:Reservation)<-[:HAS_RESERVATION]-(similarAccommodation:Accommodation)
		WHERE similarAccommodation IN userAccommodations AND otherUser.id <> "422bd1e1-6d70-42cc-adc6-89a09b313c01"
		WITH userAccommodations, otherAccommodations, COLLECT(DISTINCT otherUser) AS similarUsersByReservations
		
		MATCH (similarUser:User)-[:HAS_RESERVATION]->(:Reservation)<-[:HAS_RESERVATION]-(accommodation:Accommodation)-[:HAS_RATING]->(rating:Rating)
		WHERE similarUser IN similarUsersByReservations AND accommodation IN otherAccommodations
		WITH accommodation, AVG(rating.value) AS avgRating, SUM(CASE WHEN rating.value < 3 THEN 1 ELSE 0 END) AS lowRatingCount
		WHERE avgRating >= 3 AND lowRatingCount <= 5
		AND NOT EXISTS {
		  MATCH (accommodation)-[:HAS_RATING]->(rating:Rating)
		  WHERE rating.value < 3 AND datetime(rating.date) >= datetime() - duration({months: 3})
		}
		RETURN accommodation
		ORDER BY avgRating DESC
		LIMIT 10
	`

	params := map[string]interface{}{
		"userID": userID,
	}

	result, err := session.Run(ctx, query, params)
	if err != nil {
		return nil, err
	}

	// Extract the recommended accommodations from the query result
	recommendedAccommodations := make([]*recommendation.RecommendedAccommodationDTO, 0)
	for result.Next(ctx) {
		log.Println("Uslo u result for")
		record := result.Record()
		values := record.Values

		// Access the accommodation node from the Values slice
		if len(values) > 0 {
			log.Println("Values je veci od 0")
			accommodationNode := values[0]
			// Convert the Neo4j node to AccommodationDTO
			accommodationDTO, err := s.convertNodeToAccommodationDTO(accommodationNode)
			if err != nil {
				return nil, err
			}

			recommendedAccommodations = append(recommendedAccommodations, accommodationDTO)
		}
	}
	log.Println("Duzina preporucenih acc : ")
	log.Println(len(recommendedAccommodations))

	if err = result.Err(); err != nil {
		return nil, err
	}

	return recommendedAccommodations, nil
}

func (s RecommendationService) convertNodeToAccommodationDTO(node interface{}) (*recommendation.RecommendedAccommodationDTO, error) {
	// Extract the properties from the Neo4j node
	properties := node.(dbtype.Node).Props

	// Create a new AccommodationDTO and populate its fields with the extracted properties
	accommodationDTO := &recommendation.RecommendedAccommodationDTO{
		Id:        properties["id"].(string),
		Name:      properties["name"].(string),
		Benefits:  properties["benefits"].(string),
		MinGuests: int32(properties["minGuests"].(int64)),
		MaxGuests: int32(properties["maxGuests"].(int64)),
		BasePrice: properties["basePrice"].(float64),
	}
	log.Println(properties["name"].(string))

	return accommodationDTO, nil
}

func (s RecommendationService) SetDataForDb(accommodations []*accommodation.AccommodationDTO, users []*user.User, ratings []*rating.RatingDTO, reservations []*booking.BookingRequestsDTO) {
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
	err = s.insertRatings(session, ratings, ctx)
	if err != nil {
		s.RecommendationRepo.Logger.Println("Error inserting ratings:", err)
		// Handle the error accordingly
		return
	}

	err = s.insertReservations(session, reservations, ctx)
	if err != nil {
		s.RecommendationRepo.Logger.Println("Error inserting reservations:", err)
		// Handle the error accordingly
		return
	}
}

func (s RecommendationService) insertAccommodations(session neo4j.SessionWithContext, accommodations []*accommodation.AccommodationDTO, ctx context.Context) error {
	_, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (interface{}, error) {
			// Delete existing Accommodation nodes that are not in the new objects
			_, err := transaction.Run(ctx,
				"MATCH (a:Accommodation) WHERE NOT a.id IN $accommodationIds DETACH DELETE a",
				map[string]interface{}{
					"accommodationIds": getAccommodationIds(accommodations),
				})
			if err != nil {
				return nil, err
			}
			for _, acc := range accommodations {
				// Check if Accommodation node already exists
				result, err := transaction.Run(ctx,
					"MATCH (a:Accommodation) WHERE a.id = $id RETURN a",
					map[string]interface{}{
						"id": acc.Id,
					})
				if err != nil {
					return nil, err
				}
				if result.Next(ctx) {
					continue
				}

				log.Println("Max guests:")
				log.Println(acc.MaxGuests)
				log.Println("Min guests:")
				log.Println(acc.MinGuests)
				log.Println("Price:")
				log.Println(acc.BasePrice)
				result, err = transaction.Run(ctx,
					"CREATE (a:Accommodation) SET a.id = $id, a.name = $name, a.benefits = $benefits, a.minGuests = $minGuests, a.maxGuests = $maxGuests, a.basePrice = $basePrice",
					map[string]interface{}{
						"id":        acc.Id,
						"name":      acc.Name,
						"benefits":  acc.Benefits,
						"minGuests": uint32(acc.MinGuests),
						"maxGuests": uint32(acc.MaxGuests),
						"basePrice": float32(acc.BasePrice),
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
			// Delete existing User nodes that are not in the new objects
			_, err := transaction.Run(ctx,
				"MATCH (u:User) WHERE NOT u.id IN $userIds DETACH DELETE u",
				map[string]interface{}{
					"userIds": getUserIds(users),
				})
			if err != nil {
				return nil, err
			}
			for _, usr := range users {
				if usr.Role == user.UserRole_Guest {
					// Check if User node already exists
					result, err := transaction.Run(ctx,
						"MATCH (u:User) WHERE u.id = $id RETURN u",
						map[string]interface{}{
							"id": usr.Id,
						})
					if err != nil {
						return nil, err
					}
					if result.Next(ctx) {
						continue
					}

					result, err = transaction.Run(ctx,
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

func (s RecommendationService) insertRatings(session neo4j.SessionWithContext, ratings []*rating.RatingDTO, ctx context.Context) error {
	_, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (interface{}, error) {
			// Delete existing Rating nodes that are not in the new objects
			_, err := transaction.Run(ctx,
				"MATCH (r:Rating) WHERE NOT r.id IN $ratingIds DETACH DELETE r",
				map[string]interface{}{
					"ratingIds": getRatingIds(ratings),
				})
			if err != nil {
				return nil, err
			}

			for _, rtg := range ratings {
				// Check if Rating node already exists
				result, err := transaction.Run(ctx,
					"MATCH (r:Rating) WHERE r.id = $id RETURN r",
					map[string]interface{}{
						"id": rtg.Id,
					})
				if err != nil {
					return nil, err
				}
				if result.Next(ctx) {
					continue
				}

				dateString := rtg.LastModified
				parsedTime, err := time.Parse("2006-01-02 15:04:05.999999 -0700 MST", dateString)
				if err != nil {
					// Handle parsing error
					return nil, err
				}

				date := parsedTime.Format("2006-01-02")
				result, err = transaction.Run(ctx,
					"CREATE (r:Rating) SET r.id = $ratingId, r.value = $ratingValue, r.date = $ratingDate",
					map[string]interface{}{
						"ratingId":    rtg.Id,
						"ratingValue": rtg.Value,
						"ratingDate":  date,
					})
				if err != nil {
					return nil, err
				}
				if result.Err() != nil {
					return nil, result.Err()
				}

				// Link Rating node with Accommodation node
				_, err = transaction.Run(ctx,
					"MATCH (a:Accommodation), (r:Rating) WHERE a.id = $accommodationId AND r.id = $ratingId CREATE (a)-[:HAS_RATING]->(r)",
					map[string]interface{}{
						"accommodationId": rtg.TargetId,
						"ratingId":        rtg.Id,
					})
				if err != nil {
					return nil, err
				}

				// Link Rating node with User node
				_, err = transaction.Run(ctx,
					"MATCH (u:User), (r:Rating) WHERE u.id = $userId AND r.id = $ratingId CREATE (u)-[:HAS_RATING]->(r)",
					map[string]interface{}{
						"userId":   rtg.UserId,
						"ratingId": rtg.Id,
					})
				if err != nil {
					return nil, err
				}

			}

			return nil, nil
		})

	return err
}

func (s RecommendationService) insertReservations(session neo4j.SessionWithContext, reservations []*booking.BookingRequestsDTO, ctx context.Context) error {
	_, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (interface{}, error) {
			// Delete existing Rating nodes that are not in the new objects
			_, err := transaction.Run(ctx,
				"MATCH (r:Reservation) WHERE NOT r.id IN $reservationIds DETACH DELETE r",
				map[string]interface{}{
					"reservationIds": getReservationIds(reservations),
				})
			if err != nil {
				return nil, err
			}
			for _, reservation := range reservations {

				// Check if Rating node already exists
				result, err := transaction.Run(ctx,
					"MATCH (r:Reservation) WHERE r.id = $id RETURN r",
					map[string]interface{}{
						"id": reservation.Id,
					})
				if err != nil {
					return nil, err
				}
				if result.Next(ctx) {
					continue
				}

				result, err = transaction.Run(ctx,
					"CREATE (r:Reservation) SET r.id = $reservationId, r.startDate = $startDate, r.endDate = $endDate, r.numberOfGuests = $numberOfGuests,"+
						"r.numberOfCanceledReservations = $numberOfCanceledReservations",
					map[string]interface{}{
						"reservationId":                reservation.Id,
						"startDate":                    reservation.StartDate,
						"endDate":                      reservation.EndDate,
						"numberOfGuests":               reservation.NumberOfGuests,
						"numberOfCanceledReservations": reservation.NumberOfCanceledReservations,
					})
				if err != nil {
					return nil, err
				}
				if result.Err() != nil {
					return nil, result.Err()
				}

				// Link Booking node with Accommodation node
				_, err = transaction.Run(ctx,
					"MATCH (a:Accommodation), (r:Reservation) WHERE a.id = $accommodationId AND r.id = $reservationId CREATE (a)-[:HAS_RESERVATION]->(r)",
					map[string]interface{}{
						"accommodationId": reservation.AccommodationId,
						"reservationId":   reservation.Id,
					})
				if err != nil {
					return nil, err
				}

				// Link Booking node with User node
				_, err = transaction.Run(ctx,
					"MATCH (u:User), (r:Reservation) WHERE u.id = $userId AND r.id = $reservationId CREATE (u)-[:HAS_RESERVATION]->(r)",
					map[string]interface{}{
						"userId":        reservation.UserId,
						"reservationId": reservation.Id,
					})
				if err != nil {
					return nil, err
				}
			}

			return nil, nil
		})

	return err
}

func getReservationIds(reservations []*booking.BookingRequestsDTO) []string {
	ids := make([]string, len(reservations))
	for i, rsv := range reservations {
		ids[i] = rsv.Id
	}
	return ids
}

func getRatingIds(ratings []*rating.RatingDTO) []string {
	ids := make([]string, len(ratings))
	for i, rtg := range ratings {
		ids[i] = rtg.Id
	}
	return ids
}

func getAccommodationIds(accommodations []*accommodation.AccommodationDTO) []string {
	ids := make([]string, len(accommodations))
	for i, acc := range accommodations {
		ids[i] = acc.Id
	}
	return ids
}

func getUserIds(users []*user.User) []string {
	ids := make([]string, len(users))
	for i, userIter := range users {
		ids[i] = userIter.Id
	}
	return ids
}
