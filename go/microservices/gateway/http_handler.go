package main

import (
	"errors"
	"net/http"

	common "github.com/maslias/common"
	pb "github.com/maslias/common/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/maslias/microservices-gateway/gateway"
)

type handler struct {
	gateway gateway.OrdersGateway
}

func NewHandler(gateway gateway.OrdersGateway) *handler {
	return &handler{gateway}
}

func (h *handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /customers/{customerId}/orders", h.handleCreateOrder)
}

func (h *handler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerId := r.PathValue("customerId")

	var items []*pb.ItemsWithQantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateItems(items); err != nil {
		common.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	o, err := h.gateway.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerId: customerId,
		Items:      items,
	})

	rStatus := status.Convert(err)

	if rStatus != nil {

		if rStatus.Code() != codes.InvalidArgument {
			common.WriteJSONError(w, http.StatusBadRequest, rStatus.Message())
			return
		}

		common.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.WriteJSON(w, http.StatusOK, o)
}

func validateItems(items []*pb.ItemsWithQantity) error {
	if len(items) == 0 {
		return common.ErrNoItems
	}

	for _, i := range items {
		if i.Id == "" {
			return errors.New("item id is required")
		}

		if i.Quantity <= 0 {
			return errors.New("items must have valid quantity")
		}

	}

	return nil
}
