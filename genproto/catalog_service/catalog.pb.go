package catalog

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("catalog_service.proto", fileDescriptor_69979759328785bd) }

var fileDescriptor_69979759328785bd = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x4a, 0xf3, 0x50,
	0x10, 0x85, 0xff, 0x6e, 0xfe, 0xe8, 0xd8, 0xc4, 0x3a, 0x58, 0xc4, 0x2c, 0xf2, 0x08, 0xa1, 0x54,
	0x68, 0x17, 0x8a, 0x60, 0xa2, 0x04, 0xc1, 0x95, 0xe2, 0x5a, 0x62, 0x3b, 0xd4, 0xd0, 0xca, 0x8d,
	0xc9, 0x55, 0xc8, 0x9b, 0x08, 0xbe, 0x90, 0x4b, 0x1f, 0x41, 0xe2, 0x8b, 0x48, 0x32, 0xc9, 0x4d,
	0x9b, 0xc4, 0x6a, 0x77, 0x99, 0x73, 0xcf, 0x61, 0xce, 0x7c, 0x10, 0xe8, 0x4f, 0x7c, 0xe9, 0x2f,
	0xc4, 0xec, 0x2e, 0xa6, 0xe8, 0x25, 0x98, 0x90, 0x1d, 0x46, 0x42, 0x0a, 0xd4, 0x0a, 0xd9, 0xd4,
	0x8b, 0x0f, 0xd6, 0x87, 0x6f, 0x1a, 0x18, 0x2e, 0x2b, 0x37, 0x1c, 0xc0, 0x01, 0x74, 0xdd, 0x88,
	0x7c, 0x49, 0x67, 0xcf, 0xf2, 0x41, 0x44, 0xb8, 0x6b, 0x97, 0x11, 0x16, 0xcc, 0xba, 0x80, 0x36,
	0x6c, 0x7b, 0x24, 0x8b, 0xa1, 0xa7, 0x5e, 0x9d, 0xe4, 0x72, 0x7a, 0x4d, 0x4f, 0x4d, 0xff, 0x18,
	0x40, 0xf9, 0xe3, 0xa5, 0xc0, 0x55, 0x10, 0xcb, 0x2c, 0x70, 0x50, 0x0b, 0xb0, 0x1e, 0x87, 0x59,
	0xb5, 0xdb, 0x70, 0xba, 0x49, 0xb5, 0x01, 0x74, 0xcf, 0x69, 0x41, 0x2a, 0xd1, 0x6c, 0x67, 0x28,
	0xe5, 0xe2, 0x31, 0x94, 0x09, 0x8e, 0xc0, 0xe0, 0xf3, 0x5d, 0x5f, 0xd2, 0x4c, 0x44, 0x09, 0xee,
	0x29, 0x47, 0x29, 0x99, 0x4d, 0x09, 0x87, 0xb0, 0xe3, 0x91, 0x54, 0x63, 0x73, 0x51, 0x4b, 0xe6,
	0x04, 0xf4, 0x2a, 0x13, 0x50, 0x1b, 0x8b, 0xc3, 0x46, 0x4a, 0xd1, 0x18, 0x81, 0xc1, 0x34, 0x36,
	0x6e, 0x6a, 0x30, 0x93, 0x35, 0x65, 0xeb, 0x54, 0x6c, 0x00, 0xa6, 0xe2, 0x08, 0x31, 0x47, 0xbd,
	0xf2, 0x0b, 0x31, 0x5f, 0xda, 0x91, 0x8d, 0x79, 0x37, 0x1b, 0x34, 0x8f, 0x64, 0x6e, 0x5e, 0x47,
	0x42, 0xf9, 0xc7, 0xb0, 0x55, 0xf8, 0x63, 0xdc, 0x5f, 0x79, 0x2e, 0x41, 0xf4, 0x5b, 0xd4, 0x7c,
	0x11, 0x30, 0x84, 0x3f, 0x17, 0x03, 0x3e, 0xfe, 0x87, 0x6e, 0xf5, 0xc3, 0x4f, 0x01, 0xab, 0xc3,
	0x15, 0xb0, 0xd5, 0x32, 0x2d, 0xb0, 0xd5, 0xbe, 0x63, 0xc0, 0x6a, 0xdf, 0x6f, 0xf9, 0xda, 0x72,
	0xa7, 0xf7, 0x9e, 0x5a, 0x9d, 0x8f, 0xd4, 0xea, 0x7c, 0xa6, 0x56, 0xe7, 0xf5, 0xcb, 0xfa, 0x77,
	0xff, 0x3f, 0xff, 0x6d, 0x8f, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x80, 0x60, 0xcd, 0xd2, 0xe7,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CatalogServiceClient is the client API for CatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatalogServiceClient interface {
	CreateAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Author, error)
	GetAuthor(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Author, error)
	GetAuthors(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*AuthorListResp, error)
	UpdateAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Author, error)
	DeleteAuthor(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Empty, error)
	CreateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error)
	GetCategory(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Category, error)
	GetCategories(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*CategoryListResp, error)
	UpdateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error)
	DeleteCategory(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Empty, error)
	CreateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*BookResp, error)
	GetBook(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*BookResp, error)
	GetBooks(ctx context.Context, in *BookListReq, opts ...grpc.CallOption) (*BookListResp, error)
	UpdateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*BookResp, error)
	DeleteBook(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Empty, error)
	CreateBookCategory(ctx context.Context, in *BookCategory, opts ...grpc.CallOption) (*BookResp, error)
	DeleteBookCategory(ctx context.Context, in *BookCategory, opts ...grpc.CallOption) (*Empty, error)
}

type catalogServiceClient struct {
	cc *grpc.ClientConn
}

func NewCatalogServiceClient(cc *grpc.ClientConn) CatalogServiceClient {
	return &catalogServiceClient{cc}
}

func (c *catalogServiceClient) CreateAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/CreateAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetAuthor(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/GetAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetAuthors(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*AuthorListResp, error) {
	out := new(AuthorListResp)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/GetAuthors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/UpdateAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteAuthor(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/DeleteAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetCategory(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/GetCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetCategories(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*CategoryListResp, error) {
	out := new(CategoryListResp)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/GetCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/UpdateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteCategory(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/DeleteCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*BookResp, error) {
	out := new(BookResp)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetBook(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*BookResp, error) {
	out := new(BookResp)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetBooks(ctx context.Context, in *BookListReq, opts ...grpc.CallOption) (*BookListResp, error) {
	out := new(BookListResp)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/GetBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*BookResp, error) {
	out := new(BookResp)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteBook(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateBookCategory(ctx context.Context, in *BookCategory, opts ...grpc.CallOption) (*BookResp, error) {
	out := new(BookResp)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/CreateBookCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteBookCategory(ctx context.Context, in *BookCategory, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/DeleteBookCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServiceServer is the server API for CatalogService service.
type CatalogServiceServer interface {
	CreateAuthor(context.Context, *Author) (*Author, error)
	GetAuthor(context.Context, *ByIdReq) (*Author, error)
	GetAuthors(context.Context, *ListReq) (*AuthorListResp, error)
	UpdateAuthor(context.Context, *Author) (*Author, error)
	DeleteAuthor(context.Context, *ByIdReq) (*Empty, error)
	CreateCategory(context.Context, *Category) (*Category, error)
	GetCategory(context.Context, *ByIdReq) (*Category, error)
	GetCategories(context.Context, *ListReq) (*CategoryListResp, error)
	UpdateCategory(context.Context, *Category) (*Category, error)
	DeleteCategory(context.Context, *ByIdReq) (*Empty, error)
	CreateBook(context.Context, *Book) (*BookResp, error)
	GetBook(context.Context, *ByIdReq) (*BookResp, error)
	GetBooks(context.Context, *BookListReq) (*BookListResp, error)
	UpdateBook(context.Context, *Book) (*BookResp, error)
	DeleteBook(context.Context, *ByIdReq) (*Empty, error)
	CreateBookCategory(context.Context, *BookCategory) (*BookResp, error)
	DeleteBookCategory(context.Context, *BookCategory) (*Empty, error)
}

// UnimplementedCatalogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCatalogServiceServer struct {
}

func (*UnimplementedCatalogServiceServer) CreateAuthor(ctx context.Context, req *Author) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthor not implemented")
}
func (*UnimplementedCatalogServiceServer) GetAuthor(ctx context.Context, req *ByIdReq) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthor not implemented")
}
func (*UnimplementedCatalogServiceServer) GetAuthors(ctx context.Context, req *ListReq) (*AuthorListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthors not implemented")
}
func (*UnimplementedCatalogServiceServer) UpdateAuthor(ctx context.Context, req *Author) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthor not implemented")
}
func (*UnimplementedCatalogServiceServer) DeleteAuthor(ctx context.Context, req *ByIdReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuthor not implemented")
}
func (*UnimplementedCatalogServiceServer) CreateCategory(ctx context.Context, req *Category) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (*UnimplementedCatalogServiceServer) GetCategory(ctx context.Context, req *ByIdReq) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (*UnimplementedCatalogServiceServer) GetCategories(ctx context.Context, req *ListReq) (*CategoryListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (*UnimplementedCatalogServiceServer) UpdateCategory(ctx context.Context, req *Category) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (*UnimplementedCatalogServiceServer) DeleteCategory(ctx context.Context, req *ByIdReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (*UnimplementedCatalogServiceServer) CreateBook(ctx context.Context, req *Book) (*BookResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (*UnimplementedCatalogServiceServer) GetBook(ctx context.Context, req *ByIdReq) (*BookResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (*UnimplementedCatalogServiceServer) GetBooks(ctx context.Context, req *BookListReq) (*BookListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooks not implemented")
}
func (*UnimplementedCatalogServiceServer) UpdateBook(ctx context.Context, req *Book) (*BookResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (*UnimplementedCatalogServiceServer) DeleteBook(ctx context.Context, req *ByIdReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (*UnimplementedCatalogServiceServer) CreateBookCategory(ctx context.Context, req *BookCategory) (*BookResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBookCategory not implemented")
}
func (*UnimplementedCatalogServiceServer) DeleteBookCategory(ctx context.Context, req *BookCategory) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBookCategory not implemented")
}

func RegisterCatalogServiceServer(s *grpc.Server, srv CatalogServiceServer) {
	s.RegisterService(&_CatalogService_serviceDesc, srv)
}

func _CatalogService_CreateAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Author)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/CreateAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateAuthor(ctx, req.(*Author))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/GetAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetAuthor(ctx, req.(*ByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetAuthors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetAuthors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/GetAuthors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetAuthors(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Author)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/UpdateAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateAuthor(ctx, req.(*Author))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/DeleteAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteAuthor(ctx, req.(*ByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateCategory(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/GetCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetCategory(ctx, req.(*ByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/GetCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetCategories(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/UpdateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateCategory(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/DeleteCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteCategory(ctx, req.(*ByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetBook(ctx, req.(*ByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/GetBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetBooks(ctx, req.(*BookListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteBook(ctx, req.(*ByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateBookCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateBookCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/CreateBookCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateBookCategory(ctx, req.(*BookCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteBookCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteBookCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/DeleteBookCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteBookCategory(ctx, req.(*BookCategory))
	}
	return interceptor(ctx, in, info, handler)
}

var _CatalogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "catalog.CatalogService",
	HandlerType: (*CatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAuthor",
			Handler:    _CatalogService_CreateAuthor_Handler,
		},
		{
			MethodName: "GetAuthor",
			Handler:    _CatalogService_GetAuthor_Handler,
		},
		{
			MethodName: "GetAuthors",
			Handler:    _CatalogService_GetAuthors_Handler,
		},
		{
			MethodName: "UpdateAuthor",
			Handler:    _CatalogService_UpdateAuthor_Handler,
		},
		{
			MethodName: "DeleteAuthor",
			Handler:    _CatalogService_DeleteAuthor_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _CatalogService_CreateCategory_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _CatalogService_GetCategory_Handler,
		},
		{
			MethodName: "GetCategories",
			Handler:    _CatalogService_GetCategories_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _CatalogService_UpdateCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _CatalogService_DeleteCategory_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _CatalogService_CreateBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _CatalogService_GetBook_Handler,
		},
		{
			MethodName: "GetBooks",
			Handler:    _CatalogService_GetBooks_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _CatalogService_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _CatalogService_DeleteBook_Handler,
		},
		{
			MethodName: "CreateBookCategory",
			Handler:    _CatalogService_CreateBookCategory_Handler,
		},
		{
			MethodName: "DeleteBookCategory",
			Handler:    _CatalogService_DeleteBookCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalog_service.proto",
}