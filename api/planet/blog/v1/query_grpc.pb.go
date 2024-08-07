// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: planet/blog/v1/query.proto

package blogv1

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

const (
	Query_Params_FullMethodName          = "/planet.blog.v1.Query/Params"
	Query_GetPost_FullMethodName         = "/planet.blog.v1.Query/GetPost"
	Query_ListPost_FullMethodName        = "/planet.blog.v1.Query/ListPost"
	Query_GetSentPost_FullMethodName     = "/planet.blog.v1.Query/GetSentPost"
	Query_ListSentPost_FullMethodName    = "/planet.blog.v1.Query/ListSentPost"
	Query_GetTimeoutPost_FullMethodName  = "/planet.blog.v1.Query/GetTimeoutPost"
	Query_ListTimeoutPost_FullMethodName = "/planet.blog.v1.Query/ListTimeoutPost"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Post items.
	GetPost(ctx context.Context, in *QueryGetPostRequest, opts ...grpc.CallOption) (*QueryGetPostResponse, error)
	ListPost(ctx context.Context, in *QueryAllPostRequest, opts ...grpc.CallOption) (*QueryAllPostResponse, error)
	// Queries a list of SentPost items.
	GetSentPost(ctx context.Context, in *QueryGetSentPostRequest, opts ...grpc.CallOption) (*QueryGetSentPostResponse, error)
	ListSentPost(ctx context.Context, in *QueryAllSentPostRequest, opts ...grpc.CallOption) (*QueryAllSentPostResponse, error)
	// Queries a list of TimeoutPost items.
	GetTimeoutPost(ctx context.Context, in *QueryGetTimeoutPostRequest, opts ...grpc.CallOption) (*QueryGetTimeoutPostResponse, error)
	ListTimeoutPost(ctx context.Context, in *QueryAllTimeoutPostRequest, opts ...grpc.CallOption) (*QueryAllTimeoutPostResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetPost(ctx context.Context, in *QueryGetPostRequest, opts ...grpc.CallOption) (*QueryGetPostResponse, error) {
	out := new(QueryGetPostResponse)
	err := c.cc.Invoke(ctx, Query_GetPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListPost(ctx context.Context, in *QueryAllPostRequest, opts ...grpc.CallOption) (*QueryAllPostResponse, error) {
	out := new(QueryAllPostResponse)
	err := c.cc.Invoke(ctx, Query_ListPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetSentPost(ctx context.Context, in *QueryGetSentPostRequest, opts ...grpc.CallOption) (*QueryGetSentPostResponse, error) {
	out := new(QueryGetSentPostResponse)
	err := c.cc.Invoke(ctx, Query_GetSentPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListSentPost(ctx context.Context, in *QueryAllSentPostRequest, opts ...grpc.CallOption) (*QueryAllSentPostResponse, error) {
	out := new(QueryAllSentPostResponse)
	err := c.cc.Invoke(ctx, Query_ListSentPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetTimeoutPost(ctx context.Context, in *QueryGetTimeoutPostRequest, opts ...grpc.CallOption) (*QueryGetTimeoutPostResponse, error) {
	out := new(QueryGetTimeoutPostResponse)
	err := c.cc.Invoke(ctx, Query_GetTimeoutPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListTimeoutPost(ctx context.Context, in *QueryAllTimeoutPostRequest, opts ...grpc.CallOption) (*QueryAllTimeoutPostResponse, error) {
	out := new(QueryAllTimeoutPostResponse)
	err := c.cc.Invoke(ctx, Query_ListTimeoutPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Post items.
	GetPost(context.Context, *QueryGetPostRequest) (*QueryGetPostResponse, error)
	ListPost(context.Context, *QueryAllPostRequest) (*QueryAllPostResponse, error)
	// Queries a list of SentPost items.
	GetSentPost(context.Context, *QueryGetSentPostRequest) (*QueryGetSentPostResponse, error)
	ListSentPost(context.Context, *QueryAllSentPostRequest) (*QueryAllSentPostResponse, error)
	// Queries a list of TimeoutPost items.
	GetTimeoutPost(context.Context, *QueryGetTimeoutPostRequest) (*QueryGetTimeoutPostResponse, error)
	ListTimeoutPost(context.Context, *QueryAllTimeoutPostRequest) (*QueryAllTimeoutPostResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) GetPost(context.Context, *QueryGetPostRequest) (*QueryGetPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedQueryServer) ListPost(context.Context, *QueryAllPostRequest) (*QueryAllPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPost not implemented")
}
func (UnimplementedQueryServer) GetSentPost(context.Context, *QueryGetSentPostRequest) (*QueryGetSentPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSentPost not implemented")
}
func (UnimplementedQueryServer) ListSentPost(context.Context, *QueryAllSentPostRequest) (*QueryAllSentPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSentPost not implemented")
}
func (UnimplementedQueryServer) GetTimeoutPost(context.Context, *QueryGetTimeoutPostRequest) (*QueryGetTimeoutPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeoutPost not implemented")
}
func (UnimplementedQueryServer) ListTimeoutPost(context.Context, *QueryAllTimeoutPostRequest) (*QueryAllTimeoutPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTimeoutPost not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetPost(ctx, req.(*QueryGetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListPost(ctx, req.(*QueryAllPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetSentPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetSentPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetSentPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetSentPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetSentPost(ctx, req.(*QueryGetSentPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListSentPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllSentPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListSentPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListSentPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListSentPost(ctx, req.(*QueryAllSentPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetTimeoutPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetTimeoutPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetTimeoutPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetTimeoutPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetTimeoutPost(ctx, req.(*QueryGetTimeoutPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListTimeoutPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllTimeoutPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListTimeoutPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListTimeoutPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListTimeoutPost(ctx, req.(*QueryAllTimeoutPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "planet.blog.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _Query_GetPost_Handler,
		},
		{
			MethodName: "ListPost",
			Handler:    _Query_ListPost_Handler,
		},
		{
			MethodName: "GetSentPost",
			Handler:    _Query_GetSentPost_Handler,
		},
		{
			MethodName: "ListSentPost",
			Handler:    _Query_ListSentPost_Handler,
		},
		{
			MethodName: "GetTimeoutPost",
			Handler:    _Query_GetTimeoutPost_Handler,
		},
		{
			MethodName: "ListTimeoutPost",
			Handler:    _Query_ListTimeoutPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "planet/blog/v1/query.proto",
}
