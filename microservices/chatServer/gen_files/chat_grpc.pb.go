// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: microservices/chatServer/gen_files/chat.proto

// export PATH="$PATH:$(go env GOPATH)/bin"
// go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
// protoc --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative microservices/chatServer/gen_files/chat.proto

package chat

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
	BotChat_Recieve_FullMethodName = "/chat.BotChat/Recieve"
)

// BotChatClient is the client API for BotChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BotChatClient interface {
	// From server to bot
	// rpc Send (Message) returns (Status) {}
	// From api to bot
	Recieve(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error)
}

type botChatClient struct {
	cc grpc.ClientConnInterface
}

func NewBotChatClient(cc grpc.ClientConnInterface) BotChatClient {
	return &botChatClient{cc}
}

func (c *botChatClient) Recieve(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, BotChat_Recieve_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotChatServer is the server API for BotChat service.
// All implementations must embed UnimplementedBotChatServer
// for forward compatibility
type BotChatServer interface {
	// From server to bot
	// rpc Send (Message) returns (Status) {}
	// From api to bot
	Recieve(context.Context, *Message) (*Status, error)
	mustEmbedUnimplementedBotChatServer()
}

// UnimplementedBotChatServer must be embedded to have forward compatible implementations.
type UnimplementedBotChatServer struct {
}

func (UnimplementedBotChatServer) Recieve(context.Context, *Message) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recieve not implemented")
}
func (UnimplementedBotChatServer) mustEmbedUnimplementedBotChatServer() {}

// UnsafeBotChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BotChatServer will
// result in compilation errors.
type UnsafeBotChatServer interface {
	mustEmbedUnimplementedBotChatServer()
}

func RegisterBotChatServer(s grpc.ServiceRegistrar, srv BotChatServer) {
	s.RegisterService(&BotChat_ServiceDesc, srv)
}

func _BotChat_Recieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotChatServer).Recieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BotChat_Recieve_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotChatServer).Recieve(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// BotChat_ServiceDesc is the grpc.ServiceDesc for BotChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BotChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.BotChat",
	HandlerType: (*BotChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recieve",
			Handler:    _BotChat_Recieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "microservices/chatServer/gen_files/chat.proto",
}
