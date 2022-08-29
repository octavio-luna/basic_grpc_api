package procedures

import (
	"context"
	"errors"
	"log"

	pb "github.com/octavio-luna/basic_grpc_api/proto"
	db "github.com/octavio-luna/basic_grpc_api/server/db"
)

const (
	cacheNoExpiration = -1
)

func newString(s string) *string {
	return &s
}

type Server struct {
	pb.UnimplementedECommerceServiceServer
	handler *db.Handler
	cache   *Cache
}

func NewServer() *Server {
	return &Server{
		handler: db.New(),
		cache:   NewCache(),
	}
}

func (s *Server) AddProductToCart(c context.Context, AddProduct *pb.AddProduct) (*pb.Reply, error) {
	log.Printf("AddProductToCart: product: %d User: %d \n", AddProduct.ProductID, AddProduct.UserID)
	cart, present := s.cache.cache.Get(string(AddProduct.UserID))
	if !present {
		log.Printf("AddProductToCart: cart not found for user: %d. Creating now \n", AddProduct.UserID)
		s.cache.cache.Set(string(AddProduct.UserID), []*pb.OrderedProduct{}, cacheNoExpiration)
		cart, _ = s.cache.cache.Get(string(AddProduct.UserID))
	}

	//Check if product is already in cart and update quantity
	cart = cart.([]*pb.OrderedProduct)
	for _, product := range cart.([]*pb.OrderedProduct) {
		if product.Id == AddProduct.ProductID {
			product.Quantity += AddProduct.Quantity
			s.cache.cache.Set(string(AddProduct.UserID), cart, cacheNoExpiration)
			return &pb.Reply{
				Success: true,
			}, nil
		}
	}

	//Add product to cart
	cart = append(cart.([]*pb.OrderedProduct), &pb.OrderedProduct{
		Id:       AddProduct.ProductID,
		Quantity: AddProduct.Quantity,
	})

	s.cache.cache.Set(string(AddProduct.UserID), cart.([]*pb.OrderedProduct), cacheNoExpiration)
	return &pb.Reply{
		Success: true,
	}, nil
}

func (s *Server) GetProductsFromCart(c context.Context, user *pb.UserId) (*pb.OrderedProductList, error) {
	log.Printf("GetProductsFromCart: user: %d \n", user.Id)
	cart, present := s.cache.cache.Get(string(user.Id))
	if !present {
		return &pb.OrderedProductList{}, nil
	}
	return &pb.OrderedProductList{
		Products: cart.([]*pb.OrderedProduct),
	}, nil
}

func (s *Server) AddOrderFromCart(c context.Context, user *pb.UserId) (*pb.OrderResponse, error) {
	log.Printf("AddOrderFromCart: user: %d \n", user.Id)
	cart, err := s.GetProductsFromCart(c, user)
	if err != nil {
		log.Printf("AddOrderFromCart: error getting cart for user %d: %s \n", user.Id, err.Error())
		return &pb.OrderResponse{}, err
	}
	if len(cart.Products) == 0 {
		return &pb.OrderResponse{}, errors.New("Cart is empty")
	}

	order, err := s.CreateOrder(c, &pb.OrderRequest{
		Customerid: user.Id,
		Products:   cart.Products,
	})

	if err != nil {
		log.Printf("AddOrderFromCart: error creating order for user %d: %s \n", user.Id, err.Error())
		return &pb.OrderResponse{}, err
	}

	log.Printf("AddOrderFromCart: order created for user %d: %d. Deleting Cart \n", user.Id, order.Id)
	s.cache.cache.Delete(string(user.Id))

	return &pb.OrderResponse{
		Id:         order.Id,
		Customerid: user.Id,
		Products:   cart.Products,
	}, nil
}

func (s *Server) CreateOrder(c context.Context, orderRequest *pb.OrderRequest) (*pb.OrderResponse, error) {
	orderResponse, err := s.handler.CreateOrder(orderRequest)

	if err != nil {
		log.Printf("CreateOrder: error creating order for user %d: %s \n", orderRequest.Customerid, err.Error())
		return &pb.OrderResponse{}, err
	}
	return orderResponse, nil
}

func (s *Server) GetProductsFromOrder(c context.Context, order *pb.OrderID) (*pb.OrderedProductList, error) {
	log.Printf("GetProductsFromOrder: order: %d \n", order.Id)
	products, err := s.handler.GetProductsFromOrder(order)
	if err != nil {
		log.Println("GetProductsFromOrder: error getting products from order: ", err.Error())
		return &pb.OrderedProductList{}, err
	}
	return products, nil
}

func (s *Server) GetProductsFromIDs(c context.Context, products *pb.OrderedProductList) (*pb.ProductList, error) {
	ids := make([]int32, 0)
	for _, product := range products.Products {
		ids = append(ids, product.Id)
	}
	if len(ids) == 0 {
		return &pb.ProductList{}, nil
	}
	log.Println("GetProductsFromIDs: ids: ", ids)
	prod, err := s.handler.FindProducts(ids)
	if err != nil {
		log.Println("GetProductsFromIDs: error getting products from ids: ", err.Error())
		return nil, err
	}
	return &pb.ProductList{
		Products: prod,
	}, nil

}
