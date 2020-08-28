// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: product.proto

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

type ProductArea int32

const (
	ProductArea_A ProductArea = 0
	ProductArea_B ProductArea = 1
	ProductArea_C ProductArea = 2
)

// Enum value maps for ProductArea.
var (
	ProductArea_name = map[int32]string{
		0: "A",
		1: "B",
		2: "C",
	}
	ProductArea_value = map[string]int32{
		"A": 0,
		"B": 1,
		"C": 2,
	}
)

func (x ProductArea) Enum() *ProductArea {
	p := new(ProductArea)
	*p = x
	return p
}

func (x ProductArea) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductArea) Descriptor() protoreflect.EnumDescriptor {
	return file_product_proto_enumTypes[0].Descriptor()
}

func (ProductArea) Type() protoreflect.EnumType {
	return &file_product_proto_enumTypes[0]
}

func (x ProductArea) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductArea.Descriptor instead.
func (ProductArea) EnumDescriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

type ProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdId int32       `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"` //商品id
	Area   ProductArea `protobuf:"varint,2,opt,name=area,proto3,enum=services.ProductArea" json:"area,omitempty"`
}

func (x *ProductReq) Reset() {
	*x = ProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductReq) ProtoMessage() {}

func (x *ProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductReq.ProtoReflect.Descriptor instead.
func (*ProductReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductReq) GetProdId() int32 {
	if x != nil {
		return x.ProdId
	}
	return 0
}

func (x *ProductReq) GetArea() ProductArea {
	if x != nil {
		return x.Area
	}
	return ProductArea_A
}

type ProductResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdStock int32 `protobuf:"varint,1,opt,name=prod_stock,json=prodStock,proto3" json:"prod_stock,omitempty"` //商品库存
}

func (x *ProductResp) Reset() {
	*x = ProductResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResp) ProtoMessage() {}

func (x *ProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResp.ProtoReflect.Descriptor instead.
func (*ProductResp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductResp) GetProdStock() int32 {
	if x != nil {
		return x.ProdStock
	}
	return 0
}

type ProductsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuerySize int32 `protobuf:"varint,1,opt,name=QuerySize,proto3" json:"QuerySize,omitempty"`
}

func (x *ProductsReq) Reset() {
	*x = ProductsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsReq) ProtoMessage() {}

func (x *ProductsReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsReq.ProtoReflect.Descriptor instead.
func (*ProductsReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

func (x *ProductsReq) GetQuerySize() int32 {
	if x != nil {
		return x.QuerySize
	}
	return 0
}

type ProductsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*ProductResp `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *ProductsResp) Reset() {
	*x = ProductsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsResp) ProtoMessage() {}

func (x *ProductsResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsResp.ProtoReflect.Descriptor instead.
func (*ProductsResp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{3}
}

func (x *ProductsResp) GetProducts() []*ProductResp {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x29,
	0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41,
	0x72, 0x65, 0x61, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x22, 0x2c, 0x0a, 0x0b, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x2b, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x41, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2a, 0x22, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x41, 0x72, 0x65, 0x61, 0x12, 0x05, 0x0a, 0x01, 0x41, 0x10, 0x00, 0x12, 0x05, 0x0a,
	0x01, 0x42, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x43, 0x10, 0x02, 0x32, 0xcd, 0x01, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x15, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3d, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_product_proto_goTypes = []interface{}{
	(ProductArea)(0),     // 0: services.ProductArea
	(*ProductReq)(nil),   // 1: services.ProductReq
	(*ProductResp)(nil),  // 2: services.ProductResp
	(*ProductsReq)(nil),  // 3: services.ProductsReq
	(*ProductsResp)(nil), // 4: services.ProductsResp
	(*ProductInfo)(nil),  // 5: services.ProductInfo
}
var file_product_proto_depIdxs = []int32{
	0, // 0: services.ProductReq.area:type_name -> services.ProductArea
	2, // 1: services.ProductsResp.products:type_name -> services.ProductResp
	1, // 2: services.ProductService.GetProductStock:input_type -> services.ProductReq
	3, // 3: services.ProductService.GetProducts:input_type -> services.ProductsReq
	1, // 4: services.ProductService.GetProductInfo:input_type -> services.ProductReq
	2, // 5: services.ProductService.GetProductStock:output_type -> services.ProductResp
	4, // 6: services.ProductService.GetProducts:output_type -> services.ProductsResp
	5, // 7: services.ProductService.GetProductInfo:output_type -> services.ProductInfo
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	file_Models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductReq); i {
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
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductResp); i {
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
		file_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsReq); i {
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
		file_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsResp); i {
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
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		EnumInfos:         file_product_proto_enumTypes,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	GetProductStock(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductResp, error)
	GetProducts(ctx context.Context, in *ProductsReq, opts ...grpc.CallOption) (*ProductsResp, error)
	GetProductInfo(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductInfo, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetProductStock(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductResp, error) {
	out := new(ProductResp)
	err := c.cc.Invoke(ctx, "/services.ProductService/GetProductStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProducts(ctx context.Context, in *ProductsReq, opts ...grpc.CallOption) (*ProductsResp, error) {
	out := new(ProductsResp)
	err := c.cc.Invoke(ctx, "/services.ProductService/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProductInfo(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductInfo, error) {
	out := new(ProductInfo)
	err := c.cc.Invoke(ctx, "/services.ProductService/GetProductInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	GetProductStock(context.Context, *ProductReq) (*ProductResp, error)
	GetProducts(context.Context, *ProductsReq) (*ProductsResp, error)
	GetProductInfo(context.Context, *ProductReq) (*ProductInfo, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) GetProductStock(context.Context, *ProductReq) (*ProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductStock not implemented")
}
func (*UnimplementedProductServiceServer) GetProducts(context.Context, *ProductsReq) (*ProductsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (*UnimplementedProductServiceServer) GetProductInfo(context.Context, *ProductReq) (*ProductInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductInfo not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_GetProductStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProductService/GetProductStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductStock(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProductService/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProducts(ctx, req.(*ProductsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProductInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProductService/GetProductInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductInfo(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductStock",
			Handler:    _ProductService_GetProductStock_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _ProductService_GetProducts_Handler,
		},
		{
			MethodName: "GetProductInfo",
			Handler:    _ProductService_GetProductInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
