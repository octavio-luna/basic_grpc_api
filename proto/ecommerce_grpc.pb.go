// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: proto/ecommerce.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ECommerceServiceClient is the client API for ECommerceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ECommerceServiceClient interface {
	AddProductToCart(ctx context.Context, in *AddProduct, opts ...grpc.CallOption) (*Reply, error)
	GetProductsFromCart(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*OrderedProductList, error)
	AddOrderFromCart(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*OrderResponse, error)
	GetProductsFromOrder(ctx context.Context, in *OrderID, opts ...grpc.CallOption) (*OrderedProductList, error)
	//Added two extra methods for working with FullText products instead onf only IDs
	AddOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	GetProductsFromIDs(ctx context.Context, in *OrderedProductList, opts ...grpc.CallOption) (*ProductList, error)
}

type eCommerceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewECommerceServiceClient(cc grpc.ClientConnInterface) ECommerceServiceClient {
	return &eCommerceServiceClient{cc}
}

func (c *eCommerceServiceClient) AddProductToCart(ctx context.Context, in *AddProduct, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/grpc.ECommerceService/AddProductToCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eCommerceServiceClient) GetProductsFromCart(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*OrderedProductList, error) {
	out := new(OrderedProductList)
	err := c.cc.Invoke(ctx, "/grpc.ECommerceService/GetProductsFromCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eCommerceServiceClient) AddOrderFromCart(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/grpc.ECommerceService/AddOrderFromCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eCommerceServiceClient) GetProductsFromOrder(ctx context.Context, in *OrderID, opts ...grpc.CallOption) (*OrderedProductList, error) {
	out := new(OrderedProductList)
	err := c.cc.Invoke(ctx, "/grpc.ECommerceService/GetProductsFromOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eCommerceServiceClient) AddOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/grpc.ECommerceService/AddOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eCommerceServiceClient) GetProductsFromIDs(ctx context.Context, in *OrderedProductList, opts ...grpc.CallOption) (*ProductList, error) {
	out := new(ProductList)
	err := c.cc.Invoke(ctx, "/grpc.ECommerceService/GetProductsFromIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ECommerceServiceServer is the server API for ECommerceService service.
// All implementations must embed UnimplementedECommerceServiceServer
// for forward compatibility
type ECommerceServiceServer interface {
	AddProductToCart(context.Context, *AddProduct) (*Reply, error)
	GetProductsFromCart(context.Context, *UserId) (*OrderedProductList, error)
	AddOrderFromCart(context.Context, *UserId) (*OrderResponse, error)
	GetProductsFromOrder(context.Context, *OrderID) (*OrderedProductList, error)
	//Added two extra methods for working with FullText products instead onf only IDs
	AddOrder(context.Context, *OrderRequest) (*OrderResponse, error)
	GetProductsFromIDs(context.Context, *OrderedProductList) (*ProductList, error)
	mustEmbedUnimplementedECommerceServiceServer()
}

// UnimplementedECommerceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedECommerceServiceServer struct {
}

func (UnimplementedECommerceServiceServer) AddProductToCart(context.Context, *AddProduct) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProductToCart not implemented")
}
func (UnimplementedECommerceServiceServer) GetProductsFromCart(context.Context, *UserId) (*OrderedProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsFromCart not implemented")
}
func (UnimplementedECommerceServiceServer) AddOrderFromCart(context.Context, *UserId) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrderFromCart not implemented")
}
func (UnimplementedECommerceServiceServer) GetProductsFromOrder(context.Context, *OrderID) (*OrderedProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsFromOrder not implemented")
}
func (UnimplementedECommerceServiceServer) AddOrder(context.Context, *OrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrder not implemented")
}
func (UnimplementedECommerceServiceServer) GetProductsFromIDs(context.Context, *OrderedProductList) (*ProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsFromIDs not implemented")
}
func (UnimplementedECommerceServiceServer) mustEmbedUnimplementedECommerceServiceServer() {}

// UnsafeECommerceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ECommerceServiceServer will
// result in compilation errors.
type UnsafeECommerceServiceServer interface {
	mustEmbedUnimplementedECommerceServiceServer()
}

func RegisterECommerceServiceServer(s grpc.ServiceRegistrar, srv ECommerceServiceServer) {
	s.RegisterService(&ECommerceService_ServiceDesc, srv)
}

func _ECommerceService_AddProductToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProduct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ECommerceServiceServer).AddProductToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ECommerceService/AddProductToCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ECommerceServiceServer).AddProductToCart(ctx, req.(*AddProduct))
	}
	return interceptor(ctx, in, info, handler)
}

func _ECommerceService_GetProductsFromCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ECommerceServiceServer).GetProductsFromCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ECommerceService/GetProductsFromCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ECommerceServiceServer).GetProductsFromCart(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ECommerceService_AddOrderFromCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ECommerceServiceServer).AddOrderFromCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ECommerceService/AddOrderFromCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ECommerceServiceServer).AddOrderFromCart(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ECommerceService_GetProductsFromOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ECommerceServiceServer).GetProductsFromOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ECommerceService/GetProductsFromOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ECommerceServiceServer).GetProductsFromOrder(ctx, req.(*OrderID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ECommerceService_AddOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ECommerceServiceServer).AddOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ECommerceService/AddOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ECommerceServiceServer).AddOrder(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ECommerceService_GetProductsFromIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderedProductList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ECommerceServiceServer).GetProductsFromIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ECommerceService/GetProductsFromIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ECommerceServiceServer).GetProductsFromIDs(ctx, req.(*OrderedProductList))
	}
	return interceptor(ctx, in, info, handler)
}

// ECommerceService_ServiceDesc is the grpc.ServiceDesc for ECommerceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ECommerceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.ECommerceService",
	HandlerType: (*ECommerceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProductToCart",
			Handler:    _ECommerceService_AddProductToCart_Handler,
		},
		{
			MethodName: "GetProductsFromCart",
			Handler:    _ECommerceService_GetProductsFromCart_Handler,
		},
		{
			MethodName: "AddOrderFromCart",
			Handler:    _ECommerceService_AddOrderFromCart_Handler,
		},
		{
			MethodName: "GetProductsFromOrder",
			Handler:    _ECommerceService_GetProductsFromOrder_Handler,
		},
		{
			MethodName: "AddOrder",
			Handler:    _ECommerceService_AddOrder_Handler,
		},
		{
			MethodName: "GetProductsFromIDs",
			Handler:    _ECommerceService_GetProductsFromIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ecommerce.proto",
}
