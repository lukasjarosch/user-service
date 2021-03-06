// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

/*
Package go_micro_srv_auth is a generated protocol buffer package.

It is generated from these files:
	proto/auth/auth.proto

It has these top-level messages:
	User
	Request
	Response
	Token
	Error
*/
package go_micro_srv_auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Company  string `protobuf:"bytes,3,opt,name=company" json:"company,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
	Token    string `protobuf:"bytes,6,opt,name=token" json:"token,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Request struct {
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Response struct {
	User   *User    `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Users  []*User  `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
	Token  *Token   `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Response) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type Token struct {
	Token  string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Valid  bool     `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code        int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.auth.User")
	proto.RegisterType((*Request)(nil), "go.micro.srv.auth.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.auth.Response")
	proto.RegisterType((*Token)(nil), "go.micro.srv.auth.Token")
	proto.RegisterType((*Error)(nil), "go.micro.srv.auth.Error")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AuthService service

type AuthServiceClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
}

type authServiceClient struct {
	c           client.Client
	serviceName string
}

func NewAuthServiceClient(serviceName string, c client.Client) AuthServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.auth"
	}
	return &authServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authServiceClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AuthService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AuthService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AuthService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "AuthService.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "AuthService.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthService service

type AuthServiceHandler interface {
	Create(context.Context, *User, *Response) error
	Get(context.Context, *User, *Response) error
	GetAll(context.Context, *Request, *Response) error
	Auth(context.Context, *User, *Token) error
	ValidateToken(context.Context, *Token, *Token) error
}

func RegisterAuthServiceHandler(s server.Server, hdlr AuthServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AuthService{hdlr}, opts...))
}

type AuthService struct {
	AuthServiceHandler
}

func (h *AuthService) Create(ctx context.Context, in *User, out *Response) error {
	return h.AuthServiceHandler.Create(ctx, in, out)
}

func (h *AuthService) Get(ctx context.Context, in *User, out *Response) error {
	return h.AuthServiceHandler.Get(ctx, in, out)
}

func (h *AuthService) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.AuthServiceHandler.GetAll(ctx, in, out)
}

func (h *AuthService) Auth(ctx context.Context, in *User, out *Token) error {
	return h.AuthServiceHandler.Auth(ctx, in, out)
}

func (h *AuthService) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.AuthServiceHandler.ValidateToken(ctx, in, out)
}

func init() { proto.RegisterFile("proto/auth/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0xdd, 0x7c, 0xee, 0x6c, 0x0d, 0x0a, 0x16, 0x8a, 0xcd, 0x78, 0x19, 0x72, 0x12, 0xc4, 0x28,
	0xeb, 0x51, 0x16, 0x5c, 0x86, 0x61, 0xee, 0xf1, 0xe3, 0xde, 0x26, 0x85, 0xd3, 0x98, 0xa4, 0x63,
	0x77, 0x67, 0xc4, 0xff, 0xe0, 0x0f, 0xf3, 0x1f, 0xf8, 0x77, 0xa4, 0x2b, 0x13, 0x1d, 0x30, 0x19,
	0x71, 0x2f, 0x49, 0xbf, 0x57, 0xaf, 0x3a, 0xf5, 0xfa, 0xa5, 0xe1, 0x51, 0x67, 0xb4, 0xd3, 0x2f,
	0x64, 0xef, 0xf6, 0xfc, 0xc8, 0x19, 0xe3, 0x83, 0x4f, 0x3a, 0x6f, 0x54, 0x69, 0x74, 0x6e, 0xcd,
	0x21, 0xf7, 0x85, 0xec, 0x7b, 0x00, 0xf1, 0x7b, 0x4b, 0x06, 0xef, 0x43, 0xa8, 0x2a, 0x11, 0xac,
	0x83, 0xa7, 0x57, 0x45, 0xa8, 0x2a, 0x44, 0x88, 0x5b, 0xd9, 0x90, 0x08, 0x99, 0xe1, 0x35, 0x0a,
	0xb8, 0x2c, 0x75, 0xd3, 0xc9, 0xf6, 0x9b, 0x88, 0x98, 0x1e, 0x21, 0x3e, 0x84, 0x84, 0x1a, 0xa9,
	0x6a, 0x11, 0x33, 0x3f, 0x00, 0x5c, 0xc1, 0xa2, 0x93, 0xd6, 0x7e, 0xd5, 0xa6, 0x12, 0x09, 0x17,
	0x7e, 0x63, 0xdf, 0xe1, 0xf4, 0x67, 0x6a, 0x45, 0x3a, 0x74, 0x30, 0xc8, 0xae, 0xe0, 0xb2, 0xa0,
	0x2f, 0x3d, 0x59, 0x97, 0xfd, 0x08, 0x60, 0x51, 0x90, 0xed, 0x74, 0x6b, 0x09, 0x9f, 0x41, 0xdc,
	0x5b, 0x32, 0x3c, 0xdf, 0xf2, 0xfa, 0x71, 0xfe, 0x97, 0x91, 0xdc, 0x9b, 0x28, 0x58, 0x84, 0xcf,
	0x21, 0xf1, 0x6f, 0x2b, 0xc2, 0x75, 0x74, 0x4e, 0x3d, 0xa8, 0xf0, 0x25, 0xa4, 0x64, 0x8c, 0x36,
	0x56, 0x44, 0xac, 0x17, 0x13, 0xfa, 0xad, 0x17, 0x14, 0x47, 0x1d, 0xe6, 0xe3, 0xec, 0x31, 0x8f,
	0x33, 0xd5, 0xf0, 0xce, 0xd7, 0x47, 0x57, 0x04, 0x09, 0xe3, 0x3f, 0xa6, 0x83, 0x13, 0xd3, 0x9e,
	0x3d, 0xc8, 0x5a, 0x55, 0x7c, 0xd6, 0x8b, 0x62, 0x00, 0xff, 0x3f, 0x56, 0x76, 0x03, 0x09, 0x13,
	0x3e, 0xbb, 0x52, 0x57, 0xc4, 0x5f, 0x49, 0x0a, 0x5e, 0xe3, 0x1a, 0x96, 0x15, 0xd9, 0xd2, 0xa8,
	0xce, 0x29, 0xdd, 0x1e, 0x63, 0x3d, 0xa5, 0xae, 0x7f, 0x86, 0xb0, 0xbc, 0xed, 0xdd, 0xfe, 0x2d,
	0x99, 0x83, 0x2a, 0x09, 0xdf, 0x40, 0xba, 0x31, 0x24, 0x1d, 0xe1, 0xdc, 0x09, 0xae, 0x9e, 0x4c,
	0x14, 0xc6, 0xcc, 0xb2, 0x0b, 0xbc, 0x81, 0x68, 0x47, 0xee, 0xce, 0xed, 0x1b, 0x48, 0x77, 0xe4,
	0x6e, 0xeb, 0x1a, 0x57, 0x93, 0x42, 0xfe, 0x4f, 0xfe, 0xb5, 0xc9, 0x6b, 0x88, 0xbd, 0xa9, 0xf9,
	0x21, 0x66, 0xd3, 0xcb, 0x2e, 0x70, 0x0b, 0xf7, 0x3e, 0xf8, 0x30, 0xa4, 0xa3, 0x21, 0xc0, 0x59,
	0xf1, 0xb9, 0x6d, 0x3e, 0xa6, 0x7c, 0xfd, 0x5e, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xbf,
	0x61, 0x01, 0x97, 0x03, 0x00, 0x00,
}
