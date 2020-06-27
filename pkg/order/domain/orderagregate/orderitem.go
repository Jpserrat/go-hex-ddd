package orderagregate

//OrderItem is ad object value of the order agregate root
type OrderItem struct {
	Name  string  `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
}
