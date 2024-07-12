package mongosh

import (
	"context"
	"service/config"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client     mongo.Client
	Collection mongo.Collection
}

func GetCollection(cfg config.Config) (*MongoDB, error) {

	// url := fmt.Sprintf("mongodb://%s:%s", cfg.Mongosh.Host, cfg.Mongosh.Port)
	url := "mongodb+srv://Bek10022006:Bek10022006@fornt.otm6nho.mongodb.net/?retryWrites=true&w=majority&appName=Fornt"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, err
	}

	collection := client.Database(cfg.Mongosh.Database).Collection(cfg.Mongosh.Collection)
	return &MongoDB{Client: *client, Collection: *collection}, nil
}
