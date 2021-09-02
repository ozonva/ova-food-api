// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ova_food_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OvaFoodApiClient is the client API for OvaFoodApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OvaFoodApiClient interface {
	// Создание сущности
	CreateFoodV1(ctx context.Context, in *CreateFoodV1Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Возвращает опимание сущности пищи по её Id
	DescribeFoodV1(ctx context.Context, in *DescribeFoodV1Request, opts ...grpc.CallOption) (*DescribeFoodV1Response, error)
	//Возвращает лист хранимых сущностей пищи по списку ids
	ListFoodsV1(ctx context.Context, in *ListFoodsV1Request, opts ...grpc.CallOption) (*ListFoodsV1Response, error)
	//Возвращает страницу хранимых сущностей пищи по limit,offset
	PageFoods(ctx context.Context, in *PageFoodsV1Request, opts ...grpc.CallOption) (*PageFoodsV1Response, error)
	//Обновляет информацию о сущности пищи
	UpdateFoodV1(ctx context.Context, in *UpdateFoodV1Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//Удаляет сущность пищи по её Id
	RemoveFoodV1(ctx context.Context, in *RemoveFoodV1Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type ovaFoodApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOvaFoodApiClient(cc grpc.ClientConnInterface) OvaFoodApiClient {
	return &ovaFoodApiClient{cc}
}

func (c *ovaFoodApiClient) CreateFoodV1(ctx context.Context, in *CreateFoodV1Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ova.food.api.OvaFoodApi/CreateFoodV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ovaFoodApiClient) DescribeFoodV1(ctx context.Context, in *DescribeFoodV1Request, opts ...grpc.CallOption) (*DescribeFoodV1Response, error) {
	out := new(DescribeFoodV1Response)
	err := c.cc.Invoke(ctx, "/ova.food.api.OvaFoodApi/DescribeFoodV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ovaFoodApiClient) ListFoodsV1(ctx context.Context, in *ListFoodsV1Request, opts ...grpc.CallOption) (*ListFoodsV1Response, error) {
	out := new(ListFoodsV1Response)
	err := c.cc.Invoke(ctx, "/ova.food.api.OvaFoodApi/ListFoodsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ovaFoodApiClient) PageFoods(ctx context.Context, in *PageFoodsV1Request, opts ...grpc.CallOption) (*PageFoodsV1Response, error) {
	out := new(PageFoodsV1Response)
	err := c.cc.Invoke(ctx, "/ova.food.api.OvaFoodApi/PageFoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ovaFoodApiClient) UpdateFoodV1(ctx context.Context, in *UpdateFoodV1Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ova.food.api.OvaFoodApi/UpdateFoodV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ovaFoodApiClient) RemoveFoodV1(ctx context.Context, in *RemoveFoodV1Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ova.food.api.OvaFoodApi/RemoveFoodV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OvaFoodApiServer is the server API for OvaFoodApi service.
// All implementations must embed UnimplementedOvaFoodApiServer
// for forward compatibility
type OvaFoodApiServer interface {
	// Создание сущности
	CreateFoodV1(context.Context, *CreateFoodV1Request) (*emptypb.Empty, error)
	// Возвращает опимание сущности пищи по её Id
	DescribeFoodV1(context.Context, *DescribeFoodV1Request) (*DescribeFoodV1Response, error)
	//Возвращает лист хранимых сущностей пищи по списку ids
	ListFoodsV1(context.Context, *ListFoodsV1Request) (*ListFoodsV1Response, error)
	//Возвращает страницу хранимых сущностей пищи по limit,offset
	PageFoods(context.Context, *PageFoodsV1Request) (*PageFoodsV1Response, error)
	//Обновляет информацию о сущности пищи
	UpdateFoodV1(context.Context, *UpdateFoodV1Request) (*emptypb.Empty, error)
	//Удаляет сущность пищи по её Id
	RemoveFoodV1(context.Context, *RemoveFoodV1Request) (*emptypb.Empty, error)
	mustEmbedUnimplementedOvaFoodApiServer()
}

// UnimplementedOvaFoodApiServer must be embedded to have forward compatible implementations.
type UnimplementedOvaFoodApiServer struct {
}

func (UnimplementedOvaFoodApiServer) CreateFoodV1(context.Context, *CreateFoodV1Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFoodV1 not implemented")
}
func (UnimplementedOvaFoodApiServer) DescribeFoodV1(context.Context, *DescribeFoodV1Request) (*DescribeFoodV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeFoodV1 not implemented")
}
func (UnimplementedOvaFoodApiServer) ListFoodsV1(context.Context, *ListFoodsV1Request) (*ListFoodsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFoodsV1 not implemented")
}
func (UnimplementedOvaFoodApiServer) PageFoods(context.Context, *PageFoodsV1Request) (*PageFoodsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PageFoods not implemented")
}
func (UnimplementedOvaFoodApiServer) UpdateFoodV1(context.Context, *UpdateFoodV1Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFoodV1 not implemented")
}
func (UnimplementedOvaFoodApiServer) RemoveFoodV1(context.Context, *RemoveFoodV1Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFoodV1 not implemented")
}
func (UnimplementedOvaFoodApiServer) mustEmbedUnimplementedOvaFoodApiServer() {}

// UnsafeOvaFoodApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OvaFoodApiServer will
// result in compilation errors.
type UnsafeOvaFoodApiServer interface {
	mustEmbedUnimplementedOvaFoodApiServer()
}

func RegisterOvaFoodApiServer(s grpc.ServiceRegistrar, srv OvaFoodApiServer) {
	s.RegisterService(&OvaFoodApi_ServiceDesc, srv)
}

func _OvaFoodApi_CreateFoodV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFoodV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaFoodApiServer).CreateFoodV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.food.api.OvaFoodApi/CreateFoodV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaFoodApiServer).CreateFoodV1(ctx, req.(*CreateFoodV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OvaFoodApi_DescribeFoodV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeFoodV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaFoodApiServer).DescribeFoodV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.food.api.OvaFoodApi/DescribeFoodV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaFoodApiServer).DescribeFoodV1(ctx, req.(*DescribeFoodV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OvaFoodApi_ListFoodsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFoodsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaFoodApiServer).ListFoodsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.food.api.OvaFoodApi/ListFoodsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaFoodApiServer).ListFoodsV1(ctx, req.(*ListFoodsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OvaFoodApi_PageFoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageFoodsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaFoodApiServer).PageFoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.food.api.OvaFoodApi/PageFoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaFoodApiServer).PageFoods(ctx, req.(*PageFoodsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OvaFoodApi_UpdateFoodV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFoodV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaFoodApiServer).UpdateFoodV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.food.api.OvaFoodApi/UpdateFoodV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaFoodApiServer).UpdateFoodV1(ctx, req.(*UpdateFoodV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OvaFoodApi_RemoveFoodV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFoodV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaFoodApiServer).RemoveFoodV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.food.api.OvaFoodApi/RemoveFoodV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaFoodApiServer).RemoveFoodV1(ctx, req.(*RemoveFoodV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OvaFoodApi_ServiceDesc is the grpc.ServiceDesc for OvaFoodApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OvaFoodApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ova.food.api.OvaFoodApi",
	HandlerType: (*OvaFoodApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFoodV1",
			Handler:    _OvaFoodApi_CreateFoodV1_Handler,
		},
		{
			MethodName: "DescribeFoodV1",
			Handler:    _OvaFoodApi_DescribeFoodV1_Handler,
		},
		{
			MethodName: "ListFoodsV1",
			Handler:    _OvaFoodApi_ListFoodsV1_Handler,
		},
		{
			MethodName: "PageFoods",
			Handler:    _OvaFoodApi_PageFoods_Handler,
		},
		{
			MethodName: "UpdateFoodV1",
			Handler:    _OvaFoodApi_UpdateFoodV1_Handler,
		},
		{
			MethodName: "RemoveFoodV1",
			Handler:    _OvaFoodApi_RemoveFoodV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ova-food-api/ova-food-api.proto",
}
