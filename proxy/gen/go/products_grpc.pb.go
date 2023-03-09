// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package products

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

// ProductsServiceClient is the client API for ProductsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductsServiceClient interface {
	GetProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*Product, error)
	GetAllProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProductList, error)
	GetProductsInCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*ProductList, error)
	GetAllCategories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CategoryList, error)
}

type productsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductsServiceClient(cc grpc.ClientConnInterface) ProductsServiceClient {
	return &productsServiceClient{cc}
}

func (c *productsServiceClient) GetProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/products.ProductsService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsServiceClient) GetAllProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProductList, error) {
	out := new(ProductList)
	err := c.cc.Invoke(ctx, "/products.ProductsService/GetAllProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsServiceClient) GetProductsInCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*ProductList, error) {
	out := new(ProductList)
	err := c.cc.Invoke(ctx, "/products.ProductsService/GetProductsInCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsServiceClient) GetAllCategories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CategoryList, error) {
	out := new(CategoryList)
	err := c.cc.Invoke(ctx, "/products.ProductsService/GetAllCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductsServiceServer is the server API for ProductsService service.
// All implementations must embed UnimplementedProductsServiceServer
// for forward compatibility
type ProductsServiceServer interface {
	GetProduct(context.Context, *ProductRequest) (*Product, error)
	GetAllProducts(context.Context, *Empty) (*ProductList, error)
	GetProductsInCategory(context.Context, *CategoryRequest) (*ProductList, error)
	GetAllCategories(context.Context, *Empty) (*CategoryList, error)
	mustEmbedUnimplementedProductsServiceServer()
}

// UnimplementedProductsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductsServiceServer struct {
}

func (UnimplementedProductsServiceServer) GetProduct(context.Context, *ProductRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductsServiceServer) GetAllProducts(context.Context, *Empty) (*ProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProducts not implemented")
}
func (UnimplementedProductsServiceServer) GetProductsInCategory(context.Context, *CategoryRequest) (*ProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsInCategory not implemented")
}
func (UnimplementedProductsServiceServer) GetAllCategories(context.Context, *Empty) (*CategoryList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategories not implemented")
}
func (UnimplementedProductsServiceServer) mustEmbedUnimplementedProductsServiceServer() {}

// UnsafeProductsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductsServiceServer will
// result in compilation errors.
type UnsafeProductsServiceServer interface {
	mustEmbedUnimplementedProductsServiceServer()
}

func RegisterProductsServiceServer(s grpc.ServiceRegistrar, srv ProductsServiceServer) {
	s.RegisterService(&ProductsService_ServiceDesc, srv)
}

func _ProductsService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/products.ProductsService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServiceServer).GetProduct(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsService_GetAllProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServiceServer).GetAllProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/products.ProductsService/GetAllProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServiceServer).GetAllProducts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsService_GetProductsInCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServiceServer).GetProductsInCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/products.ProductsService/GetProductsInCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServiceServer).GetProductsInCategory(ctx, req.(*CategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsService_GetAllCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServiceServer).GetAllCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/products.ProductsService/GetAllCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServiceServer).GetAllCategories(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductsService_ServiceDesc is the grpc.ServiceDesc for ProductsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "products.ProductsService",
	HandlerType: (*ProductsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _ProductsService_GetProduct_Handler,
		},
		{
			MethodName: "GetAllProducts",
			Handler:    _ProductsService_GetAllProducts_Handler,
		},
		{
			MethodName: "GetProductsInCategory",
			Handler:    _ProductsService_GetProductsInCategory_Handler,
		},
		{
			MethodName: "GetAllCategories",
			Handler:    _ProductsService_GetAllCategories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "products.proto",
}