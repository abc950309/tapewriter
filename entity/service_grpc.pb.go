// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: service.proto

package entity

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	FileGet(ctx context.Context, in *FileGetRequest, opts ...grpc.CallOption) (*FileGetReply, error)
	FileEdit(ctx context.Context, in *FileEditRequest, opts ...grpc.CallOption) (*FileEditReply, error)
	FileMkdir(ctx context.Context, in *FileMkdirRequest, opts ...grpc.CallOption) (*FileMkdirReply, error)
	FileDelete(ctx context.Context, in *FileDeleteRequest, opts ...grpc.CallOption) (*FileDeleteReply, error)
	FileListParents(ctx context.Context, in *FileListParentsRequest, opts ...grpc.CallOption) (*FileListParentsReply, error)
	TapeMGet(ctx context.Context, in *TapeMGetRequest, opts ...grpc.CallOption) (*TapeMGetReply, error)
	JobList(ctx context.Context, in *JobListRequest, opts ...grpc.CallOption) (*JobListReply, error)
	JobCreate(ctx context.Context, in *JobCreateRequest, opts ...grpc.CallOption) (*JobCreateReply, error)
	JobNext(ctx context.Context, in *JobNextRequest, opts ...grpc.CallOption) (*JobNextReply, error)
	JobDisplay(ctx context.Context, in *JobDisplayRequest, opts ...grpc.CallOption) (*JobDisplayReply, error)
	JobGetLog(ctx context.Context, in *JobGetLogRequest, opts ...grpc.CallOption) (*JobGetLogReply, error)
	SourceList(ctx context.Context, in *SourceListRequest, opts ...grpc.CallOption) (*SourceListReply, error)
	DeviceList(ctx context.Context, in *DeviceListRequest, opts ...grpc.CallOption) (*DeviceListReply, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) FileGet(ctx context.Context, in *FileGetRequest, opts ...grpc.CallOption) (*FileGetReply, error) {
	out := new(FileGetReply)
	err := c.cc.Invoke(ctx, "/service.Service/FileGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FileEdit(ctx context.Context, in *FileEditRequest, opts ...grpc.CallOption) (*FileEditReply, error) {
	out := new(FileEditReply)
	err := c.cc.Invoke(ctx, "/service.Service/FileEdit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FileMkdir(ctx context.Context, in *FileMkdirRequest, opts ...grpc.CallOption) (*FileMkdirReply, error) {
	out := new(FileMkdirReply)
	err := c.cc.Invoke(ctx, "/service.Service/FileMkdir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FileDelete(ctx context.Context, in *FileDeleteRequest, opts ...grpc.CallOption) (*FileDeleteReply, error) {
	out := new(FileDeleteReply)
	err := c.cc.Invoke(ctx, "/service.Service/FileDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FileListParents(ctx context.Context, in *FileListParentsRequest, opts ...grpc.CallOption) (*FileListParentsReply, error) {
	out := new(FileListParentsReply)
	err := c.cc.Invoke(ctx, "/service.Service/FileListParents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) TapeMGet(ctx context.Context, in *TapeMGetRequest, opts ...grpc.CallOption) (*TapeMGetReply, error) {
	out := new(TapeMGetReply)
	err := c.cc.Invoke(ctx, "/service.Service/TapeMGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) JobList(ctx context.Context, in *JobListRequest, opts ...grpc.CallOption) (*JobListReply, error) {
	out := new(JobListReply)
	err := c.cc.Invoke(ctx, "/service.Service/JobList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) JobCreate(ctx context.Context, in *JobCreateRequest, opts ...grpc.CallOption) (*JobCreateReply, error) {
	out := new(JobCreateReply)
	err := c.cc.Invoke(ctx, "/service.Service/JobCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) JobNext(ctx context.Context, in *JobNextRequest, opts ...grpc.CallOption) (*JobNextReply, error) {
	out := new(JobNextReply)
	err := c.cc.Invoke(ctx, "/service.Service/JobNext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) JobDisplay(ctx context.Context, in *JobDisplayRequest, opts ...grpc.CallOption) (*JobDisplayReply, error) {
	out := new(JobDisplayReply)
	err := c.cc.Invoke(ctx, "/service.Service/JobDisplay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) JobGetLog(ctx context.Context, in *JobGetLogRequest, opts ...grpc.CallOption) (*JobGetLogReply, error) {
	out := new(JobGetLogReply)
	err := c.cc.Invoke(ctx, "/service.Service/JobGetLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SourceList(ctx context.Context, in *SourceListRequest, opts ...grpc.CallOption) (*SourceListReply, error) {
	out := new(SourceListReply)
	err := c.cc.Invoke(ctx, "/service.Service/SourceList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeviceList(ctx context.Context, in *DeviceListRequest, opts ...grpc.CallOption) (*DeviceListReply, error) {
	out := new(DeviceListReply)
	err := c.cc.Invoke(ctx, "/service.Service/DeviceList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	FileGet(context.Context, *FileGetRequest) (*FileGetReply, error)
	FileEdit(context.Context, *FileEditRequest) (*FileEditReply, error)
	FileMkdir(context.Context, *FileMkdirRequest) (*FileMkdirReply, error)
	FileDelete(context.Context, *FileDeleteRequest) (*FileDeleteReply, error)
	FileListParents(context.Context, *FileListParentsRequest) (*FileListParentsReply, error)
	TapeMGet(context.Context, *TapeMGetRequest) (*TapeMGetReply, error)
	JobList(context.Context, *JobListRequest) (*JobListReply, error)
	JobCreate(context.Context, *JobCreateRequest) (*JobCreateReply, error)
	JobNext(context.Context, *JobNextRequest) (*JobNextReply, error)
	JobDisplay(context.Context, *JobDisplayRequest) (*JobDisplayReply, error)
	JobGetLog(context.Context, *JobGetLogRequest) (*JobGetLogReply, error)
	SourceList(context.Context, *SourceListRequest) (*SourceListReply, error)
	DeviceList(context.Context, *DeviceListRequest) (*DeviceListReply, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) FileGet(context.Context, *FileGetRequest) (*FileGetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileGet not implemented")
}
func (UnimplementedServiceServer) FileEdit(context.Context, *FileEditRequest) (*FileEditReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileEdit not implemented")
}
func (UnimplementedServiceServer) FileMkdir(context.Context, *FileMkdirRequest) (*FileMkdirReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileMkdir not implemented")
}
func (UnimplementedServiceServer) FileDelete(context.Context, *FileDeleteRequest) (*FileDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileDelete not implemented")
}
func (UnimplementedServiceServer) FileListParents(context.Context, *FileListParentsRequest) (*FileListParentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileListParents not implemented")
}
func (UnimplementedServiceServer) TapeMGet(context.Context, *TapeMGetRequest) (*TapeMGetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TapeMGet not implemented")
}
func (UnimplementedServiceServer) JobList(context.Context, *JobListRequest) (*JobListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobList not implemented")
}
func (UnimplementedServiceServer) JobCreate(context.Context, *JobCreateRequest) (*JobCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobCreate not implemented")
}
func (UnimplementedServiceServer) JobNext(context.Context, *JobNextRequest) (*JobNextReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobNext not implemented")
}
func (UnimplementedServiceServer) JobDisplay(context.Context, *JobDisplayRequest) (*JobDisplayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobDisplay not implemented")
}
func (UnimplementedServiceServer) JobGetLog(context.Context, *JobGetLogRequest) (*JobGetLogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobGetLog not implemented")
}
func (UnimplementedServiceServer) SourceList(context.Context, *SourceListRequest) (*SourceListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SourceList not implemented")
}
func (UnimplementedServiceServer) DeviceList(context.Context, *DeviceListRequest) (*DeviceListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceList not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_FileGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FileGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/FileGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FileGet(ctx, req.(*FileGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_FileEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileEditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FileEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/FileEdit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FileEdit(ctx, req.(*FileEditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_FileMkdir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileMkdirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FileMkdir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/FileMkdir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FileMkdir(ctx, req.(*FileMkdirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_FileDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FileDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/FileDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FileDelete(ctx, req.(*FileDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_FileListParents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileListParentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FileListParents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/FileListParents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FileListParents(ctx, req.(*FileListParentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_TapeMGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TapeMGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).TapeMGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/TapeMGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).TapeMGet(ctx, req.(*TapeMGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_JobList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).JobList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/JobList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).JobList(ctx, req.(*JobListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_JobCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).JobCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/JobCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).JobCreate(ctx, req.(*JobCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_JobNext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobNextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).JobNext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/JobNext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).JobNext(ctx, req.(*JobNextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_JobDisplay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobDisplayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).JobDisplay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/JobDisplay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).JobDisplay(ctx, req.(*JobDisplayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_JobGetLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobGetLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).JobGetLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/JobGetLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).JobGetLog(ctx, req.(*JobGetLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_SourceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SourceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).SourceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/SourceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).SourceList(ctx, req.(*SourceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeviceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeviceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/DeviceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeviceList(ctx, req.(*DeviceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FileGet",
			Handler:    _Service_FileGet_Handler,
		},
		{
			MethodName: "FileEdit",
			Handler:    _Service_FileEdit_Handler,
		},
		{
			MethodName: "FileMkdir",
			Handler:    _Service_FileMkdir_Handler,
		},
		{
			MethodName: "FileDelete",
			Handler:    _Service_FileDelete_Handler,
		},
		{
			MethodName: "FileListParents",
			Handler:    _Service_FileListParents_Handler,
		},
		{
			MethodName: "TapeMGet",
			Handler:    _Service_TapeMGet_Handler,
		},
		{
			MethodName: "JobList",
			Handler:    _Service_JobList_Handler,
		},
		{
			MethodName: "JobCreate",
			Handler:    _Service_JobCreate_Handler,
		},
		{
			MethodName: "JobNext",
			Handler:    _Service_JobNext_Handler,
		},
		{
			MethodName: "JobDisplay",
			Handler:    _Service_JobDisplay_Handler,
		},
		{
			MethodName: "JobGetLog",
			Handler:    _Service_JobGetLog_Handler,
		},
		{
			MethodName: "SourceList",
			Handler:    _Service_SourceList_Handler,
		},
		{
			MethodName: "DeviceList",
			Handler:    _Service_DeviceList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
