// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.4
// source: scanner/api/v1/orchestrator_scan_service.proto

package scannerV1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	OrchestratorScanService_GetKubeVulnerabilities_FullMethodName      = "/scannerV1.OrchestratorScanService/GetKubeVulnerabilities"
	OrchestratorScanService_GetOpenShiftVulnerabilities_FullMethodName = "/scannerV1.OrchestratorScanService/GetOpenShiftVulnerabilities"
	OrchestratorScanService_GetIstioVulnerabilities_FullMethodName     = "/scannerV1.OrchestratorScanService/GetIstioVulnerabilities"
)

// OrchestratorScanServiceClient is the client API for OrchestratorScanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// OrchestratorScanService APIs can be used to get vulnerabilities for Kubernetes and Openshift components.
type OrchestratorScanServiceClient interface {
	GetKubeVulnerabilities(ctx context.Context, in *GetKubeVulnerabilitiesRequest, opts ...grpc.CallOption) (*GetKubeVulnerabilitiesResponse, error)
	GetOpenShiftVulnerabilities(ctx context.Context, in *GetOpenShiftVulnerabilitiesRequest, opts ...grpc.CallOption) (*GetOpenShiftVulnerabilitiesResponse, error)
	GetIstioVulnerabilities(ctx context.Context, in *GetIstioVulnerabilitiesRequest, opts ...grpc.CallOption) (*GetIstioVulnerabilitiesResponse, error)
}

type orchestratorScanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrchestratorScanServiceClient(cc grpc.ClientConnInterface) OrchestratorScanServiceClient {
	return &orchestratorScanServiceClient{cc}
}

func (c *orchestratorScanServiceClient) GetKubeVulnerabilities(ctx context.Context, in *GetKubeVulnerabilitiesRequest, opts ...grpc.CallOption) (*GetKubeVulnerabilitiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetKubeVulnerabilitiesResponse)
	err := c.cc.Invoke(ctx, OrchestratorScanService_GetKubeVulnerabilities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorScanServiceClient) GetOpenShiftVulnerabilities(ctx context.Context, in *GetOpenShiftVulnerabilitiesRequest, opts ...grpc.CallOption) (*GetOpenShiftVulnerabilitiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOpenShiftVulnerabilitiesResponse)
	err := c.cc.Invoke(ctx, OrchestratorScanService_GetOpenShiftVulnerabilities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorScanServiceClient) GetIstioVulnerabilities(ctx context.Context, in *GetIstioVulnerabilitiesRequest, opts ...grpc.CallOption) (*GetIstioVulnerabilitiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIstioVulnerabilitiesResponse)
	err := c.cc.Invoke(ctx, OrchestratorScanService_GetIstioVulnerabilities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrchestratorScanServiceServer is the server API for OrchestratorScanService service.
// All implementations should embed UnimplementedOrchestratorScanServiceServer
// for forward compatibility
//
// OrchestratorScanService APIs can be used to get vulnerabilities for Kubernetes and Openshift components.
type OrchestratorScanServiceServer interface {
	GetKubeVulnerabilities(context.Context, *GetKubeVulnerabilitiesRequest) (*GetKubeVulnerabilitiesResponse, error)
	GetOpenShiftVulnerabilities(context.Context, *GetOpenShiftVulnerabilitiesRequest) (*GetOpenShiftVulnerabilitiesResponse, error)
	GetIstioVulnerabilities(context.Context, *GetIstioVulnerabilitiesRequest) (*GetIstioVulnerabilitiesResponse, error)
}

// UnimplementedOrchestratorScanServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrchestratorScanServiceServer struct {
}

func (UnimplementedOrchestratorScanServiceServer) GetKubeVulnerabilities(context.Context, *GetKubeVulnerabilitiesRequest) (*GetKubeVulnerabilitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKubeVulnerabilities not implemented")
}
func (UnimplementedOrchestratorScanServiceServer) GetOpenShiftVulnerabilities(context.Context, *GetOpenShiftVulnerabilitiesRequest) (*GetOpenShiftVulnerabilitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOpenShiftVulnerabilities not implemented")
}
func (UnimplementedOrchestratorScanServiceServer) GetIstioVulnerabilities(context.Context, *GetIstioVulnerabilitiesRequest) (*GetIstioVulnerabilitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIstioVulnerabilities not implemented")
}

// UnsafeOrchestratorScanServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrchestratorScanServiceServer will
// result in compilation errors.
type UnsafeOrchestratorScanServiceServer interface {
	mustEmbedUnimplementedOrchestratorScanServiceServer()
}

func RegisterOrchestratorScanServiceServer(s grpc.ServiceRegistrar, srv OrchestratorScanServiceServer) {
	s.RegisterService(&OrchestratorScanService_ServiceDesc, srv)
}

func _OrchestratorScanService_GetKubeVulnerabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKubeVulnerabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorScanServiceServer).GetKubeVulnerabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrchestratorScanService_GetKubeVulnerabilities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorScanServiceServer).GetKubeVulnerabilities(ctx, req.(*GetKubeVulnerabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrchestratorScanService_GetOpenShiftVulnerabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOpenShiftVulnerabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorScanServiceServer).GetOpenShiftVulnerabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrchestratorScanService_GetOpenShiftVulnerabilities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorScanServiceServer).GetOpenShiftVulnerabilities(ctx, req.(*GetOpenShiftVulnerabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrchestratorScanService_GetIstioVulnerabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIstioVulnerabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorScanServiceServer).GetIstioVulnerabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrchestratorScanService_GetIstioVulnerabilities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorScanServiceServer).GetIstioVulnerabilities(ctx, req.(*GetIstioVulnerabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrchestratorScanService_ServiceDesc is the grpc.ServiceDesc for OrchestratorScanService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrchestratorScanService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scannerV1.OrchestratorScanService",
	HandlerType: (*OrchestratorScanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKubeVulnerabilities",
			Handler:    _OrchestratorScanService_GetKubeVulnerabilities_Handler,
		},
		{
			MethodName: "GetOpenShiftVulnerabilities",
			Handler:    _OrchestratorScanService_GetOpenShiftVulnerabilities_Handler,
		},
		{
			MethodName: "GetIstioVulnerabilities",
			Handler:    _OrchestratorScanService_GetIstioVulnerabilities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scanner/api/v1/orchestrator_scan_service.proto",
}