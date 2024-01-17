// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: proto/operator.proto

package operator

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
	Operator_GetCapabilities_FullMethodName       = "/cnpgi.adapter.v1.Operator/GetCapabilities"
	Operator_ValidateClusterCreate_FullMethodName = "/cnpgi.adapter.v1.Operator/ValidateClusterCreate"
	Operator_ValidateClusterChange_FullMethodName = "/cnpgi.adapter.v1.Operator/ValidateClusterChange"
	Operator_MutateCluster_FullMethodName         = "/cnpgi.adapter.v1.Operator/MutateCluster"
	Operator_MutatePod_FullMethodName             = "/cnpgi.adapter.v1.Operator/MutatePod"
)

// OperatorClient is the client API for Operator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperatorClient interface {
	// GetCapabilities gets the capabilities of the WAL service
	GetCapabilities(ctx context.Context, in *OperatorCapabilitiesRequest, opts ...grpc.CallOption) (*OperatorCapabilitiesResult, error)
	// ValidateCreate improves the behaviour of the validating webhook that
	// is called on creation of the Cluster resources
	ValidateClusterCreate(ctx context.Context, in *OperatorValidateClusterCreateRequest, opts ...grpc.CallOption) (*OperatorValidateClusterCreateResult, error)
	// ValidateClusterChange improves the behavior of the validating webhook of
	// is called on updates of the Cluster resources
	ValidateClusterChange(ctx context.Context, in *OperatorValidateClusterChangeRequest, opts ...grpc.CallOption) (*OperatorValidateClusterChangeResult, error)
	// MutateCluster fills in the defaults inside a Cluster resource
	MutateCluster(ctx context.Context, in *OperatorMutateClusterRequest, opts ...grpc.CallOption) (*OperatorMutateClusterResult, error)
	// MutatePod reconciles a Pod definition before it
	// is applied in the Kubernetes cluster
	MutatePod(ctx context.Context, in *OperatorMutatePodRequest, opts ...grpc.CallOption) (*OperatorMutatePodResult, error)
}

type operatorClient struct {
	cc grpc.ClientConnInterface
}

func NewOperatorClient(cc grpc.ClientConnInterface) OperatorClient {
	return &operatorClient{cc}
}

func (c *operatorClient) GetCapabilities(ctx context.Context, in *OperatorCapabilitiesRequest, opts ...grpc.CallOption) (*OperatorCapabilitiesResult, error) {
	out := new(OperatorCapabilitiesResult)
	err := c.cc.Invoke(ctx, Operator_GetCapabilities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) ValidateClusterCreate(ctx context.Context, in *OperatorValidateClusterCreateRequest, opts ...grpc.CallOption) (*OperatorValidateClusterCreateResult, error) {
	out := new(OperatorValidateClusterCreateResult)
	err := c.cc.Invoke(ctx, Operator_ValidateClusterCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) ValidateClusterChange(ctx context.Context, in *OperatorValidateClusterChangeRequest, opts ...grpc.CallOption) (*OperatorValidateClusterChangeResult, error) {
	out := new(OperatorValidateClusterChangeResult)
	err := c.cc.Invoke(ctx, Operator_ValidateClusterChange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) MutateCluster(ctx context.Context, in *OperatorMutateClusterRequest, opts ...grpc.CallOption) (*OperatorMutateClusterResult, error) {
	out := new(OperatorMutateClusterResult)
	err := c.cc.Invoke(ctx, Operator_MutateCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) MutatePod(ctx context.Context, in *OperatorMutatePodRequest, opts ...grpc.CallOption) (*OperatorMutatePodResult, error) {
	out := new(OperatorMutatePodResult)
	err := c.cc.Invoke(ctx, Operator_MutatePod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatorServer is the server API for Operator service.
// All implementations must embed UnimplementedOperatorServer
// for forward compatibility
type OperatorServer interface {
	// GetCapabilities gets the capabilities of the WAL service
	GetCapabilities(context.Context, *OperatorCapabilitiesRequest) (*OperatorCapabilitiesResult, error)
	// ValidateCreate improves the behaviour of the validating webhook that
	// is called on creation of the Cluster resources
	ValidateClusterCreate(context.Context, *OperatorValidateClusterCreateRequest) (*OperatorValidateClusterCreateResult, error)
	// ValidateClusterChange improves the behavior of the validating webhook of
	// is called on updates of the Cluster resources
	ValidateClusterChange(context.Context, *OperatorValidateClusterChangeRequest) (*OperatorValidateClusterChangeResult, error)
	// MutateCluster fills in the defaults inside a Cluster resource
	MutateCluster(context.Context, *OperatorMutateClusterRequest) (*OperatorMutateClusterResult, error)
	// MutatePod reconciles a Pod definition before it
	// is applied in the Kubernetes cluster
	MutatePod(context.Context, *OperatorMutatePodRequest) (*OperatorMutatePodResult, error)
	mustEmbedUnimplementedOperatorServer()
}

// UnimplementedOperatorServer must be embedded to have forward compatible implementations.
type UnimplementedOperatorServer struct {
}

func (UnimplementedOperatorServer) GetCapabilities(context.Context, *OperatorCapabilitiesRequest) (*OperatorCapabilitiesResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCapabilities not implemented")
}
func (UnimplementedOperatorServer) ValidateClusterCreate(context.Context, *OperatorValidateClusterCreateRequest) (*OperatorValidateClusterCreateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateClusterCreate not implemented")
}
func (UnimplementedOperatorServer) ValidateClusterChange(context.Context, *OperatorValidateClusterChangeRequest) (*OperatorValidateClusterChangeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateClusterChange not implemented")
}
func (UnimplementedOperatorServer) MutateCluster(context.Context, *OperatorMutateClusterRequest) (*OperatorMutateClusterResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCluster not implemented")
}
func (UnimplementedOperatorServer) MutatePod(context.Context, *OperatorMutatePodRequest) (*OperatorMutatePodResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutatePod not implemented")
}
func (UnimplementedOperatorServer) mustEmbedUnimplementedOperatorServer() {}

// UnsafeOperatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperatorServer will
// result in compilation errors.
type UnsafeOperatorServer interface {
	mustEmbedUnimplementedOperatorServer()
}

func RegisterOperatorServer(s grpc.ServiceRegistrar, srv OperatorServer) {
	s.RegisterService(&Operator_ServiceDesc, srv)
}

func _Operator_GetCapabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperatorCapabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).GetCapabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Operator_GetCapabilities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).GetCapabilities(ctx, req.(*OperatorCapabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_ValidateClusterCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperatorValidateClusterCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).ValidateClusterCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Operator_ValidateClusterCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).ValidateClusterCreate(ctx, req.(*OperatorValidateClusterCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_ValidateClusterChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperatorValidateClusterChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).ValidateClusterChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Operator_ValidateClusterChange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).ValidateClusterChange(ctx, req.(*OperatorValidateClusterChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_MutateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperatorMutateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).MutateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Operator_MutateCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).MutateCluster(ctx, req.(*OperatorMutateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_MutatePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperatorMutatePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).MutatePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Operator_MutatePod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).MutatePod(ctx, req.(*OperatorMutatePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Operator_ServiceDesc is the grpc.ServiceDesc for Operator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Operator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cnpgi.adapter.v1.Operator",
	HandlerType: (*OperatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCapabilities",
			Handler:    _Operator_GetCapabilities_Handler,
		},
		{
			MethodName: "ValidateClusterCreate",
			Handler:    _Operator_ValidateClusterCreate_Handler,
		},
		{
			MethodName: "ValidateClusterChange",
			Handler:    _Operator_ValidateClusterChange_Handler,
		},
		{
			MethodName: "MutateCluster",
			Handler:    _Operator_MutateCluster_Handler,
		},
		{
			MethodName: "MutatePod",
			Handler:    _Operator_MutatePod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/operator.proto",
}