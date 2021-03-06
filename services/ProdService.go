package services

import (
	"context"
)

type ProdService struct {
	UnimplementedProdServiceServer
}

//func (this *ProdService) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error) {
//
//}

func (this *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error) {
	var stock int32 = 0
	if request.ProdArea == ProdAreas_A {
		stock = 30
	} else if request.ProdArea == ProdAreas_B {
		stock = 31
	} else {
		stock = 50
	}
	return &ProdResponse{ProdStock: stock}, nil
}

func (this *ProdService) GetProdStocks(context.Context, *QuerySize) (*ProdResponseList, error) {
	Prodres := []*ProdResponse{
		&ProdResponse{ProdStock: 28},
		&ProdResponse{ProdStock: 29},
	}
	return &ProdResponseList{
		Prodres: Prodres,
	}, nil
}

func (this *ProdService) GetProdInfo(context.Context, *ProdRequest) (*ProdModel, error) {
	ret := ProdModel{
		ProdId:    101,
		ProdName:  "测试商品",
		ProdPrice: 20.5,
	}
	return &ret, nil
}
