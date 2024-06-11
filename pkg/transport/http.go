package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"

	"gin-gokit-microservice/pkg/endpoint"
	"gin-gokit-microservice/pkg/service"
)

// SetupRoutes 设置Gin路由
func SetupRoutes(router *gin.Engine) {
	svc := service.NewService()
	eps := endpoint.MakeEndpoints(svc)

	helloHandler := httptransport.NewServer(
		eps.HelloEndpoint,
		decodeHelloRequest,
		encodeResponse,
	)

	router.POST("/hello", gin.WrapF(helloHandler.ServeHTTP))
}

func decodeHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.HelloRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
