// Code generated by protoc-gen-go-http. DO NOT EDIT.

package v1

import (
	context "context"
	middleware "github.com/go-kratos/kratos/v2/middleware"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
// context./http./middleware.
const _ = http1.SupportPackageIsVersion1

type GreeterHTTPServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterHTTPServer(s http1.ServiceRegistrar, srv GreeterHTTPServer) {
	s.RegisterService(&_HTTP_Greeter_serviceDesc, srv)
}

func _HTTP_Greeter_SayHello_0(srv interface{}, ctx context.Context, req *http.Request, dec func(interface{}) error, m middleware.Middleware) (interface{}, error) {
	var in HelloRequest

	if err := http1.BindVars(req, &in); err != nil {
		return nil, err
	}

	if err := http1.BindForm(req, &in); err != nil {
		return nil, err
	}

	h := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, &in)
	}
	out, err := m(h)(ctx, &in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _HTTP_Greeter_serviceDesc = http1.ServiceDesc{
	ServiceName: "helloworld.v1.Greeter",
	Methods: []http1.MethodDesc{

		{
			Path:    "/helloworld/{name}",
			Method:  "GET",
			Handler: _HTTP_Greeter_SayHello_0,
		},
	},
	Metadata: "api/helloworld/v1/greeter.proto",
}
