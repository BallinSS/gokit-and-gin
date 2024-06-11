package endpoint

import (
	"context"

	"gin-gokit-microservice/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

// HelloRequest 请求结构体
type HelloRequest struct {
	Name string `json:"name"`
}

// HelloResponse 响应结构体
type HelloResponse struct {
	Message string `json:"message"`
}

// Endpoints 定义所有的端点
type Endpoints struct {
	HelloEndpoint endpoint.Endpoint
}

// MakeEndpoints 创建所有端点
func MakeEndpoints(svc service.Service) Endpoints {
	return Endpoints{
		HelloEndpoint: makeHelloEndpoint(svc),
	}
}

func makeHelloEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HelloRequest)
		message, err := svc.Hello(req.Name)
		if err != nil {
			return nil, err
		}
		return HelloResponse{Message: message}, nil
	}
}
