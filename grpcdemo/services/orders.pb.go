// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: orders.proto

package services

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type OrdersResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string           `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	M    map[int32]string `protobuf:"bytes,3,rep,name=m,proto3" json:"m,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *OrdersResp) Reset() {
	*x = OrdersResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdersResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdersResp) ProtoMessage() {}

func (x *OrdersResp) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdersResp.ProtoReflect.Descriptor instead.
func (*OrdersResp) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{0}
}

func (x *OrdersResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OrdersResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *OrdersResp) GetM() map[int32]string {
	if x != nil {
		return x.M
	}
	return nil
}

var File_orders_proto protoreflect.FileDescriptor

var file_orders_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x29, 0x0a, 0x01, 0x6d,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x4d, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x01, 0x6d, 0x1a, 0x34, 0x0a, 0x06, 0x4d, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x45, 0x0a, 0x0c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x08,
	0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x14, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orders_proto_rawDescOnce sync.Once
	file_orders_proto_rawDescData = file_orders_proto_rawDesc
)

func file_orders_proto_rawDescGZIP() []byte {
	file_orders_proto_rawDescOnce.Do(func() {
		file_orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_proto_rawDescData)
	})
	return file_orders_proto_rawDescData
}

var file_orders_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_orders_proto_goTypes = []interface{}{
	(*OrdersResp)(nil), // 0: services.OrdersResp
	nil,                // 1: services.OrdersResp.MEntry
	(*OrderInfo)(nil),  // 2: services.OrderInfo
}
var file_orders_proto_depIdxs = []int32{
	1, // 0: services.OrdersResp.m:type_name -> services.OrdersResp.MEntry
	2, // 1: services.OrderService.NewOrder:input_type -> services.OrderInfo
	0, // 2: services.OrderService.NewOrder:output_type -> services.OrdersResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_orders_proto_init() }
func file_orders_proto_init() {
	if File_orders_proto != nil {
		return
	}
	file_Models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_orders_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrdersResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orders_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orders_proto_goTypes,
		DependencyIndexes: file_orders_proto_depIdxs,
		MessageInfos:      file_orders_proto_msgTypes,
	}.Build()
	File_orders_proto = out.File
	file_orders_proto_rawDesc = nil
	file_orders_proto_goTypes = nil
	file_orders_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	NewOrder(ctx context.Context, in *OrderInfo, opts ...grpc.CallOption) (*OrdersResp, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) NewOrder(ctx context.Context, in *OrderInfo, opts ...grpc.CallOption) (*OrdersResp, error) {
	out := new(OrdersResp)
	err := c.cc.Invoke(ctx, "/services.OrderService/NewOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	NewOrder(context.Context, *OrderInfo) (*OrdersResp, error)
}

// UnimplementedOrderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (*UnimplementedOrderServiceServer) NewOrder(context.Context, *OrderInfo) (*OrdersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewOrder not implemented")
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_NewOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).NewOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.OrderService/NewOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).NewOrder(ctx, req.(*OrderInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewOrder",
			Handler:    _OrderService_NewOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders.proto",
}