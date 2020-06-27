package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Jpserrat/hex-ddd-example/pkg/mongoclient"
	"github.com/Jpserrat/hex-ddd-example/pkg/order/application/orderhandlers"
	"github.com/Jpserrat/hex-ddd-example/pkg/order/domain/orderservice"
	"github.com/Jpserrat/hex-ddd-example/pkg/order/infrastructure/ordermongo"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	client, err := mongoclient.New("mongodb://localhost:27017", 10)

	if err != nil {
		fmt.Printf("mongo error: %v", err)
	}

	//order intances
	orderrep := ordermongo.New(10*time.Second, *client.Database("ddd").Collection("orders"))
	ordersvc := orderservice.New(orderrep)

	//order handler
	r.HandleFunc("/orders", orderhandlers.Create(ordersvc)).Methods("POST")
	r.HandleFunc("/orders/complete/{id}", orderhandlers.CompleteOrder(ordersvc)).Methods("POST")
	r.HandleFunc("/orders/reject/{id}", orderhandlers.RejectOrder(ordersvc)).Methods("POST")
	r.HandleFunc("/orders/accept/{id}", orderhandlers.AcceptOrder(ordersvc)).Methods("POST")

	errs := make(chan error, 2)

	go func() {
		fmt.Println("Listening on port: 8080")
		errs <- http.ListenAndServe(":8080", r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("\nTermineted %v\n", <-errs)

}
