package services

import (
	"context"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"recommendation-service/infrastructure/database"
	"recommendation-service/proto/recommendation"
)

type RecommendationService struct {
	RecommendationRepo *database.RecommendationRepo
}

//func (s *RecommendationService) Init() error {
//	// Access the Neo4j driver and logger from the RecommendationRepo
//	driver := s.RecommendationRepo.Driver
//	logger := s.RecommendationRepo.Logger
//	ctx := context.Background()
//
//	// Create the nodes and relationships
//	session := driver.NewSession(ctx, neo4j.SessionConfig{})
//	defer session.Close(ctx)
//
//	// Create schema constraints
//	createConstraints := []string{
//		"CREATE CONSTRAINT ON (u:User) ASSERT u.userId IS UNIQUE",
//		"CREATE CONSTRAINT ON (a:Accommodation) ASSERT a.accommodationId IS UNIQUE",
//		"CREATE CONSTRAINT ON (r:Reservation) ASSERT r.reservationId IS UNIQUE",
//		"CREATE CONSTRAINT ON (rt:Rating) ASSERT rt.ratingId IS UNIQUE",
//	}
//
//	for _, query := range createConstraints {
//		_, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
//			result, err := transaction.Run(ctx, query, map[string]any{"id": person.Born, "name": person.Name})
//			if err != nil {
//				return nil, err
//			}
//			return result.Consume()
//		})
//
//		if err != nil {
//			logger.Fatal(err)
//		}
//	}
//
//	logger.Println("Schema constraints created successfully")
//
//	return nil
//}

type Person struct {
	Name string
	Born string
}

func (s *RecommendationService) Test() (*recommendation.Test_Response, error) {
	var personOne = Person{Name: "Luka", Born: "16.06.1998"}
	err := writePerson(&personOne, s)

	if err != nil {
		return nil, errors.New("Error creating person")
	}
	var personTwo = Person{Name: "Pera", Born: "16.06.1998"}
	err = writePerson(&personTwo, s)

	if err != nil {
		return nil, errors.New("Error creating person")
	}

	return &recommendation.Test_Response{}, nil
}

func writePerson(person *Person, s *RecommendationService) error {
	ctx := context.Background()
	session := s.RecommendationRepo.Driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	// ExecuteWrite for write transactions (Create/Update/Delete)
	savedPerson, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"CREATE (p:Person) SET p.born = $born, p.name = $name RETURN p.name + ', from node ' + id(p)",
				map[string]any{"born": person.Born, "name": person.Name})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		})
	if err != nil {
		s.RecommendationRepo.Logger.Println("Error inserting Person:", err)
		return err
	}
	s.RecommendationRepo.Logger.Println(savedPerson.(string))
	return nil
}
