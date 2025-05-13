package excel

import (
	"context"

	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/usecases/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type orderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *orderRepository {
	return &orderRepository{
		collection: db.Collection("order"),
	}
}

func (or *orderRepository) SaveOrder(ctx context.Context, orders []domain.Order) ([]string, error) {
	var docs []interface{}

	for _, order := range orders {
		docs = append(docs, bson.M{
			"order_id":     order.OrderID,
			"customer":     order.Customer,
			"total_amount": order.TotalAmount,
			"date":         order.Date,
			"status":       order.Status,
		})
	}

	_, err := or.collection.InsertMany(ctx, docs)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
