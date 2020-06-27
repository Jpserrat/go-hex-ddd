package ordermongo

import (
	"context"
	"fmt"
	"time"

	"github.com/Jpserrat/hex-ddd-example/pkg/order/domain/orderagregate"
	"github.com/Jpserrat/hex-ddd-example/pkg/order/domain/orderrepository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ordermongo struct {
	timeout    time.Duration
	collection mongo.Collection
}

//New returns a new mongo repository for the order domain
func New(timeout time.Duration, col mongo.Collection) orderrepository.Repository {
	return &ordermongo{
		timeout:    timeout,
		collection: col,
	}
}

func (m ordermongo) Create(order *orderagregate.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	_, err := m.collection.InsertOne(ctx, order)

	if err != nil {
		return err
	}
	return nil
}

func (m ordermongo) Save(order *orderagregate.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(order.ID)

	if err != nil {
		return err
	}

	order.ID = ""

	_, err = m.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": *order})

	if err != nil {
		fmt.Printf("save error: %v", err)
		return err
	}

	return nil
}

func (m ordermongo) FindByID(id string) (*orderagregate.Order, error) {
	var order orderagregate.Order

	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return &order, err
	}

	m.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&order)

	return &order, nil
}
