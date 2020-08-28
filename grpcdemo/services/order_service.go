package services

import (
	"context"
	"fmt"
)

type OrderService struct {
}

func (o *OrderService) NewOrder(ctx context.Context, info *OrderInfo) (*OrdersResp, error) {
	fmt.Println(info)
	r
}
