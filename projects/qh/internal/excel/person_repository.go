package excel

import (
	"context"

	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/usecases/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		collection: db.Collection("persons"),
	}
}

func (r *MongoRepository) SavePerson(ctx context.Context, persons []domain.Person_2) ([]string, error) {
	var docs []interface{}

	for _, p := range persons {
		docs = append(docs, bson.M{
			"first_name": p.FirstName,
			"last_name":  p.LastName,
			"age":        p.Age,
			"phone":      p.Phone,
		})
	}

	_, err := r.collection.InsertMany(ctx, docs)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
