package service

import (
	"context"
	"github.com/fastmall/order/api"
)

type OrderService struct {
	api.UnimplementedOrderServiceServer
}

func (s OrderService) GetOrderDetail(ctx context.Context, req *api.GetOrderDetailRequest) (*api.GetOrderDetailResponse, error) {
	res := api.GetOrderDetailResponse{
		OrderId: req.OrderId,
	}

	return &res, nil
}

func (s OrderService) mustEmbedUnimplementedOrderServiceServer() {
	panic("implement me")
}
