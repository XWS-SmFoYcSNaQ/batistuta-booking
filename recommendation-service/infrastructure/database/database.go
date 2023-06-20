package database

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
	"os"
	"recommendation-service/config"
)

// PersonRepo NoSQL: PersonRepo struct encapsulating Neo4J api client
type RecommendationRepo struct {
	// Thread-safe instance which maintains a database connection pool
	Driver neo4j.DriverWithContext
	Logger *log.Logger
}

func Create(cfg config.Config, logger *log.Logger) (*RecommendationRepo, error) {
	// Local instance
	uri := os.Getenv("RECOMMENDATION_DATABASE_ADDRESS")
	user := os.Getenv("RECOMMENDATION_DATABASE_USERNAME")
	pass := os.Getenv("RECOMMENDATION_DATABASE_PASSWORD")
	auth := neo4j.BasicAuth(user, pass, "")

	driver, err := neo4j.NewDriverWithContext(uri, auth)
	if err != nil {
		logger.Panic(err)
		return nil, err
	}

	// Return repository with logger and DB session
	return &RecommendationRepo{
		Driver: driver,
		Logger: logger,
	}, nil
}

// CheckConnection Check if connection is established
func (rr *RecommendationRepo) CheckConnection() {
	ctx := context.Background()
	err := rr.Driver.VerifyConnectivity(ctx)
	if err != nil {
		rr.Logger.Panic(err)
		return
	}
	// Print Neo4J server address
	rr.Logger.Printf(`Neo4J server address: %s`, rr.Driver.Target().Host)
}

// CloseDriverConnection Disconnect from database
func (rr *RecommendationRepo) CloseDriverConnection(ctx context.Context) {
	rr.Driver.Close(ctx)
}
