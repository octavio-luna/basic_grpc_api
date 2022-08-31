package db

import (
	"log"

	pb "github.com/octavio-luna/basic_grpc_api/proto"
)

func (h *Handler) CreateOrder(orderReq *pb.OrderRequest) (*pb.OrderResponse, error) {
	type order struct {
		Id     int64
		UserId int64
	}

	log.Println("CreateOrder: user: %d products: %d", orderReq.Customerid, orderReq.Products)

	o := order{}
	err := h.db.Table("orders").Last(&o).Error
	if err != nil {
		if err.Error() != "record not found" {
			if err != nil {
				return &pb.OrderResponse{}, err
			}
		}
		o.Id = 0
	}

	err = h.db.Table("orders").Create(&order{
		Id:     o.Id + 1,
		UserId: int64(orderReq.Customerid),
	}).Error
	if err != nil {
		return &pb.OrderResponse{}, err
	}

	type orderProduct struct {
		Order_id   int64 `gorm:"column:order_id"`
		Product_id int64 `gorm:"column:product_id"`
		Quantity   int64 `gorm:"column:quantity"`
	}

	for _, p := range orderReq.Products {
		err := h.db.Table("order_products").Create(&orderProduct{
			Order_id:   o.Id + 1,
			Product_id: int64(p.Id),
			Quantity:   int64(p.Quantity),
		}).Error

		if err != nil {
			return &pb.OrderResponse{}, err
		}
	}

	return &pb.OrderResponse{
		Id:         int32(o.Id + 1),
		Customerid: int32(orderReq.Customerid),
		Products:   orderReq.Products,
	}, nil
}

func (h *Handler) GetProductsFromOrder(orderId *pb.OrderID) (*pb.OrderedProductList, error) {
	type orderProduct struct {
		Order_id   int64 `gorm:"column:order_id"`
		Product_id int64 `gorm:"column:product_id"`
		Quantity   int64 `gorm:"column:quantity"`
	}

	orderProducts := make([]orderProduct, 0)
	err := h.db.Table("order_products").Where("order_id = ?", orderId.Id).Scan(&orderProducts).Error
	if err != nil {
		return &pb.OrderedProductList{}, err
	}

	products := make([]*pb.OrderedProduct, 0)
	for _, p := range orderProducts {
		product := pb.OrderedProduct{}
		err := h.db.Table("products").Where("id = ?", p.Product_id).Scan(&product).Error
		if err != nil {
			return &pb.OrderedProductList{}, err
		}
		products = append(products, &pb.OrderedProduct{
			Id:       int32(p.Product_id),
			Quantity: int32(p.Quantity),
		})
	}
	return &pb.OrderedProductList{
		Products: products,
	}, nil
}
