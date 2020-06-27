package orderhandlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jpserrat/hex-ddd-example/pkg/order/application/orderrequest"
	"github.com/Jpserrat/hex-ddd-example/pkg/order/domain/orderservice"
	"github.com/gorilla/mux"
)

//Create handler
func Create(s orderservice.Service) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var request orderrequest.Create

		rw.Header().Set("Content-type", "application/json")

		err := json.NewDecoder(req.Body).Decode(&request)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		}

		err = s.Create(request.OrderItems)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		}

		rw.WriteHeader(http.StatusOK)
	}

}

//CompleteOrder handler
func CompleteOrder(s orderservice.Service) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-type", "application/json")

		vars := mux.Vars(req)

		err := s.CompleteOrder(vars["id"])

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		}

		rw.WriteHeader(http.StatusOK)
	}
}

//RejectOrder handler
func RejectOrder(s orderservice.Service) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-type", "application/json")

		vars := mux.Vars(req)

		err := s.RejectOrder(vars["id"])

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		}

		rw.WriteHeader(http.StatusOK)
	}
}

//AcceptOrder handler
func AcceptOrder(s orderservice.Service) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-type", "application/json")

		vars := mux.Vars(req)

		err := s.AcceptOrder(vars["id"])

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		}

		rw.WriteHeader(http.StatusOK)
	}
}
