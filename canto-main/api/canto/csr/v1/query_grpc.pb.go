// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: canto/csr/v1/query.proto

package csrv1

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
	Query_Params_FullMethodName        = "/canto.csr.v1.Query/Params"
	Query_CSRs_FullMethodName          = "/canto.csr.v1.Query/CSRs"
	Query_CSRByNFT_FullMethodName      = "/canto.csr.v1.Query/CSRByNFT"
	Query_CSRByContract_FullMethodName = "/canto.csr.v1.Query/CSRByContract"
	Query_Turnstile_FullMethodName     = "/canto.csr.v1.Query/Turnstile"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// query all registered CSRs
	CSRs(ctx context.Context, in *QueryCSRsRequest, opts ...grpc.CallOption) (*QueryCSRsResponse, error)
	// query a specific CSR by the nftId
	CSRByNFT(ctx context.Context, in *QueryCSRByNFTRequest, opts ...grpc.CallOption) (*QueryCSRByNFTResponse, error)
	// query a CSR by smart contract address
	CSRByContract(ctx context.Context, in *QueryCSRByContractRequest, opts ...grpc.CallOption) (*QueryCSRByContractResponse, error)
	// query the turnstile address
	Turnstile(ctx context.Context, in *QueryTurnstileRequest, opts ...grpc.CallOption) (*QueryTurnstileResponse, error)
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

func (c *queryClient) CSRs(ctx context.Context, in *QueryCSRsRequest, opts ...grpc.CallOption) (*QueryCSRsResponse, error) {
	out := new(QueryCSRsResponse)
	err := c.cc.Invoke(ctx, Query_CSRs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CSRByNFT(ctx context.Context, in *QueryCSRByNFTRequest, opts ...grpc.CallOption) (*QueryCSRByNFTResponse, error) {
	out := new(QueryCSRByNFTResponse)
	err := c.cc.Invoke(ctx, Query_CSRByNFT_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CSRByContract(ctx context.Context, in *QueryCSRByContractRequest, opts ...grpc.CallOption) (*QueryCSRByContractResponse, error) {
	out := new(QueryCSRByContractResponse)
	err := c.cc.Invoke(ctx, Query_CSRByContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Turnstile(ctx context.Context, in *QueryTurnstileRequest, opts ...grpc.CallOption) (*QueryTurnstileResponse, error) {
	out := new(QueryTurnstileResponse)
	err := c.cc.Invoke(ctx, Query_Turnstile_FullMethodName, in, out, opts...)
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
	// query all registered CSRs
	CSRs(context.Context, *QueryCSRsRequest) (*QueryCSRsResponse, error)
	// query a specific CSR by the nftId
	CSRByNFT(context.Context, *QueryCSRByNFTRequest) (*QueryCSRByNFTResponse, error)
	// query a CSR by smart contract address
	CSRByContract(context.Context, *QueryCSRByContractRequest) (*QueryCSRByContractResponse, error)
	// query the turnstile address
	Turnstile(context.Context, *QueryTurnstileRequest) (*QueryTurnstileResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) CSRs(context.Context, *QueryCSRsRequest) (*QueryCSRsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CSRs not implemented")
}
func (UnimplementedQueryServer) CSRByNFT(context.Context, *QueryCSRByNFTRequest) (*QueryCSRByNFTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CSRByNFT not implemented")
}
func (UnimplementedQueryServer) CSRByContract(context.Context, *QueryCSRByContractRequest) (*QueryCSRByContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CSRByContract not implemented")
}
func (UnimplementedQueryServer) Turnstile(context.Context, *QueryTurnstileRequest) (*QueryTurnstileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Turnstile not implemented")
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

func _Query_CSRs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCSRsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CSRs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_CSRs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CSRs(ctx, req.(*QueryCSRsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CSRByNFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCSRByNFTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CSRByNFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_CSRByNFT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CSRByNFT(ctx, req.(*QueryCSRByNFTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CSRByContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCSRByContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CSRByContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_CSRByContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CSRByContract(ctx, req.(*QueryCSRByContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Turnstile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTurnstileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Turnstile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Turnstile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Turnstile(ctx, req.(*QueryTurnstileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "canto.csr.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "CSRs",
			Handler:    _Query_CSRs_Handler,
		},
		{
			MethodName: "CSRByNFT",
			Handler:    _Query_CSRByNFT_Handler,
		},
		{
			MethodName: "CSRByContract",
			Handler:    _Query_CSRByContract_Handler,
		},
		{
			MethodName: "Turnstile",
			Handler:    _Query_Turnstile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "canto/csr/v1/query.proto",
}