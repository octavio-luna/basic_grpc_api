package db

import (
	"errors"

	pb "github.com/octavio-luna/basic_grpc_api/proto"
)

func (h *Handler) FindProducts(ids []int32) ([]*pb.Product, error) {
	products := make([]pb.Product, 0)
	err := h.db.Table("products").Select("*").Where("id IN (?)", ids).Scan(&products).Error
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, errors.New("No products found")
	}
	prodPointers := make([]*pb.Product, 0)
	for _, product := range products {
		prodPointers = append(prodPointers, &pb.Product{
			Id:           product.Id,
			Label:        product.Label,
			Type:         product.Type,
			Downloadlink: product.Downloadlink,
			Weight:       product.Weight,
		})
	}

	return prodPointers, nil
}
