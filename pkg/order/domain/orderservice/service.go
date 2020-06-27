package orderservice

import (
	"fmt"

	"github.com/Jpserrat/hex-ddd-example/pkg/order/domain/orderagregate"
	"github.com/Jpserrat/hex-ddd-example/pkg/order/domain/orderrepository"
)

//Service holds the interface of the order domain service
type Service interface {
	Create(orderItem []orderagregate.OrderItem) error
	CompleteOrder(id string) error
	AcceptOrder(id string) error
	RejectOrder(id string) error
}

type service struct {
	rep orderrepository.Repository
}

//New order
func New(rep orderrepository.Repository) Service {
	return &service{
		rep: rep,
	}
}

//Create service
func (s service) Create(orderItems []orderagregate.OrderItem) error {

	newOrder := orderagregate.CreateOrder(orderItems)

	if err := s.rep.Create(&newOrder); err != nil {
		return err
	}

	return nil
}

//AcceptOrder changes the order status to accept
func (s service) AcceptOrder(id string) error {
	order, err := s.rep.FindByID(id)

	if err != nil {
		return err
	}

	fmt.Println(order.Status)

	order.Accept()

	fmt.Println(order.Status)

	err = s.rep.Save(order)

	if err != nil {
		return err
	}

	return nil
}

//RejectOrder changes the order status to reject
func (s service) RejectOrder(id string) error {
	order, err := s.rep.FindByID(id)

	if err != nil {
		return err
	}

	order.Reject()

	err = s.rep.Save(order)

	if err != nil {
		return err
	}

	return nil
}

//CompleteOrder changes the order status to complete
func (s service) CompleteOrder(id string) error {
	order, err := s.rep.FindByID(id)

	if err != nil {
		return err
	}

	order.Complete()

	err = s.rep.Save(order)

	if err != nil {
		return err
	}

	return nil
}
