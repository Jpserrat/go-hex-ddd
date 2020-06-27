package orderrequest

import "github.com/Jpserrat/hex-ddd-example/pkg/order/domain/orderagregate"

//Create holds the implementation of the Create method params
type Create struct {
	OrderItems []orderagregate.OrderItem `json:"orderItems"`
}
