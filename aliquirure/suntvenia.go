import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connect to mongoDB
func connect() (*mongo.Database, error) {
	// get the mongoDB url from the environment variable
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		return nil, fmt.Errorf("MONGO_URL environment variable not set")
	}

	// set the connect timeout to 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create a new mongo client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, fmt.Errorf("mongo.Connect: %v", err)
	}

	// ping the database to check if it's alive
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("client.Ping: %v", err)
	}

	log.Println("Connected to MongoDB")

	// return the database handle
	return client.Database("test"), nil
}
  
