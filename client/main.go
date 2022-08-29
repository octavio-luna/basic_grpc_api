package main

import (
	"context"
	"fmt"

	pb "github.com/octavio-luna/basic_grpc_api/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		panic("Can't connect with server. Err: " + err.Error())
	}

	defer conn.Close()
	serviceClient := pb.NewECommerceServiceClient(conn)

	resp, err := serviceClient.AddProductToCart(context.Background(), &pb.AddProduct{
		ProductID: 1,
		Quantity:  2,
		UserID:    1,
	})

	resp, err = serviceClient.AddProductToCart(context.Background(), &pb.AddProduct{
		ProductID: 2,
		Quantity:  2,
		UserID:    1,
	})

	if err != nil {
		panic("Can't add product to cart. Err: " + err.Error())
	}
	fmt.Printf("Product added to cart. Success: %t \n", resp.Success)

	resp2, err := serviceClient.GetProductsFromCart(context.Background(), &pb.UserId{Id: 1})
	if err != nil {
		panic("Can't get products from cart. Err: " + err.Error())
	}
	fmt.Printf("Products from cart: \n")
	for _, product := range resp2.Products {
		fmt.Printf("Product: %d, Quantity: %d \n", product.Id, product.Quantity)
	}

	resp3, err := serviceClient.GetProductsFromIDs(context.Background(), resp2)
	if err != nil {
		panic("Can't get products from IDs. Err: " + err.Error())
	}
	fmt.Printf("Products from IDs: \n")
	for _, product := range resp3.Products {
		fmt.Printf("%v \n", product)
	}

	resp4, err := serviceClient.AddOrderFromCart(context.Background(), &pb.UserId{Id: 1})
	if err != nil {
		panic("Can't add order. Err: " + err.Error())
	}
	fmt.Printf("Order added. Order ID: %d \n", resp4.Id)

	resp5, err := serviceClient.GetProductsFromOrder(context.Background(), &pb.OrderID{Id: resp4.Id})
	if err != nil {
		panic("Can't get products from order. Err: " + err.Error())
	}
	fmt.Printf("Products from order: \n")
	for _, product := range resp5.Products {
		fmt.Printf("%v \n", product)
	}
}
