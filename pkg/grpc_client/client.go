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

type GrpcClient struct {
	orderService       order_service.OrderServiceClient
	connOrderProductService order_product_service.OrderProductsClient
	orderStatusService order_notes.OrderStatusNotesClient
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {

	connOrderService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.OrderServiceHost, cfg.OrderServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	connOrderProductService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.OrderServiceHost, cfg.OrderServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	connOrderStatusService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.OrderServiceHost, cfg.OrderServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	return &GrpcClient{
		orderService:       order_service.NewOrderServiceClient(connOrderService),
		connOrderProductService: order_product_service.NewOrderProductsClient(connOrderProductService),
		orderStatusService: order_notes.NewOrderStatusNotesClient(connOrderStatusService),
	}, nil
}

func (g *GrpcClient) OrderService() order_service.OrderServiceClient {
	return g.orderService
}

func (g *GrpcClient) OrderProductsService() order_product_service.OrderProductsClient {
	return g.connOrderProductService
}

func (g *GrpcClient) OrderProductNotesService() order_notes.OrderStatusNotesClient {
	return g.orderStatusService
}
