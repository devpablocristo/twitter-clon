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

func (r *MongoRepository) SavePerson(ctx context.Context, person domain.Person_2) error {
	_, err := r.collection.InsertOne(ctx, bson.M{
		"first_name": person.FirstName,
		"last_name":  person.LastName,
		"age":        person.Age,
		"phone":      person.Phone,
	})

	return err
}
