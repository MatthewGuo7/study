package services

import (
	context "context"
	"fmt"
)

type ProductService struct {
}

func (p *ProductService) GetProductStock(ctx context.Context, req *ProductReq) (*ProductResp, error) {
	fmt.Printf("req = %+v", req)
	resp := &ProductResp{ProdStock: 2}
	return resp, nil
}
