package orderrepository

import (
	"github.com/Jpserrat/hex-ddd-example/pkg/order/domain/orderagregate"
)

//Repository holds the interface for the repository of the order domain
type Repository interface {
	Create(order *orderagregate.Order) error
	Save(order *orderagregate.Order) error
	FindByID(id string) (*orderagregate.Order, error)
}
