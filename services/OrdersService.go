package services

import (
	"context"
	"fmt"
)

type OrdersService struct {
	UnimplementedOrderServiceServer
}

func (this *OrdersService) NewOrder(ctx context.Context, orderRequest *OrderRequest) (*OrderResponse, error) {
	fmt.Println(orderRequest.OrderMain.OrderDetail)
	return &OrderResponse{
		Status:  "OK",
		Message: "success",
	}, nil
}
