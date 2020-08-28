package services

import (
	context "context"
	"fmt"
)

type ProductService struct {
}

func (p *ProductService) GetProductInfo(ctx context.Context, req *ProductReq) (*ProductInfo, error) {
	fmt.Println(req)
	return &ProductInfo{
		ProductID:     50,
		ProductName:   "test",
	}, nil
}

func (p *ProductService) GetProducts(ctx context.Context, req *ProductsReq) (*ProductsResp, error) {
	fmt.Printf("req = %+v", req)
	products := []*ProductResp{
		&ProductResp{ProdStock: 1},
		&ProductResp{ProdStock: 3},
		&ProductResp{ProdStock: 4},
	}
	resp := &ProductsResp{
		Products: products,
	}

	return resp, nil
}

func (p *ProductService) GetProductStock(ctx context.Context, req *ProductReq) (*ProductResp, error) {
	fmt.Printf("req = %+v", req)
	resp := &ProductResp{ProdStock: 2}
	if req.Area == ProductArea_B {
		resp.ProdStock = 20
	}


	return resp, nil
}
