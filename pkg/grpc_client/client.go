package grpc_client

import (
	"fmt"
	"uzumapi/genproto/order_notes"
	"uzumapi/genproto/order_product_service"
	"uzumapi/genproto/order_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"uzumapi/config"
)

// GrpcClientI ...
type GrpcClientI interface {
	OrderService() order_service.OrderServiceClient
	OrderProduct() order_product_service.OrderProductsClient
	OrderNotes() order_notes.OrderStatusNotesClient
}

// GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {

	connOrder, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.OrderServiceHost, cfg.OrderServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("order service dial host: %s port:%s err: %s",
			cfg.OrderServiceHost, cfg.OrderServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"order_service": order_service.NewOrderServiceClient(connOrder),
		},
	}, nil
}

func (g *GrpcClient) OrderService() order_service.OrderServiceClient {
	return g.connections["order_service"].(order_service.OrderServiceClient)
}

func (g *GrpcClient) OrderProductsService() order_product_service.OrderProductsClient {
	return g.connections["order_product_service"].(order_product_service.OrderProductsClient)
}

func (g *GrpcClient) OrderProductNotesService() order_notes.OrderStatusNotesClient {
	return g.connections["order_notes"].(order_notes.OrderStatusNotesClient)
}
