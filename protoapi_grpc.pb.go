// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protoapi

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

// PracticeClient is the client API for Practice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PracticeClient interface {
	DoSum(ctx context.Context, in *RequestNumbers, opts ...grpc.CallOption) (*Result, error)
	DoMinus(ctx context.Context, in *RequestNumbers, opts ...grpc.CallOption) (*Result, error)
	GetStrLength(ctx context.Context, in *RequestStrLen, opts ...grpc.CallOption) (*Result, error)
}

type practiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPracticeClient(cc grpc.ClientConnInterface) PracticeClient {
	return &practiceClient{cc}
}

func (c *practiceClient) DoSum(ctx context.Context, in *RequestNumbers, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/Practice/DoSum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *practiceClient) DoMinus(ctx context.Context, in *RequestNumbers, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/Practice/DoMinus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *practiceClient) GetStrLength(ctx context.Context, in *RequestStrLen, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/Practice/GetStrLength", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PracticeServer is the server API for Practice service.
// All implementations must embed UnimplementedPracticeServer
// for forward compatibility
type PracticeServer interface {
	DoSum(context.Context, *RequestNumbers) (*Result, error)
	DoMinus(context.Context, *RequestNumbers) (*Result, error)
	GetStrLength(context.Context, *RequestStrLen) (*Result, error)
	mustEmbedUnimplementedPracticeServer()
}

// UnimplementedPracticeServer must be embedded to have forward compatible implementations.
type UnimplementedPracticeServer struct {
}

func (UnimplementedPracticeServer) DoSum(context.Context, *RequestNumbers) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSum not implemented")
}
func (UnimplementedPracticeServer) DoMinus(context.Context, *RequestNumbers) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoMinus not implemented")
}
func (UnimplementedPracticeServer) GetStrLength(context.Context, *RequestStrLen) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStrLength not implemented")
}
func (UnimplementedPracticeServer) mustEmbedUnimplementedPracticeServer() {}

// UnsafePracticeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PracticeServer will
// result in compilation errors.
type UnsafePracticeServer interface {
	mustEmbedUnimplementedPracticeServer()
}

func RegisterPracticeServer(s grpc.ServiceRegistrar, srv PracticeServer) {
	s.RegisterService(&Practice_ServiceDesc, srv)
}

func _Practice_DoSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestNumbers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PracticeServer).DoSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Practice/DoSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PracticeServer).DoSum(ctx, req.(*RequestNumbers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Practice_DoMinus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestNumbers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PracticeServer).DoMinus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Practice/DoMinus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PracticeServer).DoMinus(ctx, req.(*RequestNumbers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Practice_GetStrLength_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestStrLen)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PracticeServer).GetStrLength(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Practice/GetStrLength",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PracticeServer).GetStrLength(ctx, req.(*RequestStrLen))
	}
	return interceptor(ctx, in, info, handler)
}

// Practice_ServiceDesc is the grpc.ServiceDesc for Practice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Practice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Practice",
	HandlerType: (*PracticeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoSum",
			Handler:    _Practice_DoSum_Handler,
		},
		{
			MethodName: "DoMinus",
			Handler:    _Practice_DoMinus_Handler,
		},
		{
			MethodName: "GetStrLength",
			Handler:    _Practice_GetStrLength_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoapi.proto",
}
