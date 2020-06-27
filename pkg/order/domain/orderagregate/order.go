package orderagregate

//Order agregate root for the order domain
type Order struct {
	ID         string      `json:"id,omitempty" bson:"_id,omitempty"`
	OrderItems []OrderItem `json:"orderItems" bson:"order_items"`
	Status     string      `json:"status" bson:"status"`
}

//CreateOrder order
func CreateOrder(orderItem []OrderItem) Order {
	return Order{
		Status:     "PENDING",
		OrderItems: orderItem,
	}
}

//Complete dfajdsfad
func (o *Order) Complete() {
	o.Status = "COMPLETE"
}

//Reject changes the order status to rejected
func (o *Order) Reject() {
	o.Status = "REJECT"
}

//Accept changes the order status to accept
func (o *Order) Accept() {
	o.Status = "ACCEPT"
}
