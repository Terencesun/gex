// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package orderservice

import (
	"context"

	"github.com/luxun9527/gex/app/order/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CancelOrderReq              = pb.CancelOrderReq
	CreateOrderReq              = pb.CreateOrderReq
	FreezeUserAssetResp         = pb.FreezeUserAssetResp
	GetOrderAllPendingOrderResp = pb.GetOrderAllPendingOrderResp
	GetOrderListByUserReq       = pb.GetOrderListByUserReq
	GetOrderListByUserResp      = pb.GetOrderListByUserResp
	Order                       = pb.Order
	OrderEmpty                  = pb.OrderEmpty
	UpdateEntrustOrderReq       = pb.UpdateEntrustOrderReq
	UpdateOrderStatusReq        = pb.UpdateOrderStatusReq

	OrderService interface {
		// 下单
		Order(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*OrderEmpty, error)
		// 创建订单,下单有分布式事务要处理分为两个接口
		CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*OrderEmpty, error)
		// 获取用户订单列表
		GetOrderList(ctx context.Context, in *GetOrderListByUserReq, opts ...grpc.CallOption) (*GetOrderListByUserResp, error)
		// 取消订单
		CancelOrder(ctx context.Context, in *CancelOrderReq, opts ...grpc.CallOption) (*OrderEmpty, error)
		// 下单补偿
		CreateOrderRevert(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*OrderEmpty, error)
		// 获取所有订单状态为未成交或部分成交的订单
		GetOrderAllPendingOrder(ctx context.Context, in *OrderEmpty, opts ...grpc.CallOption) (pb.OrderService_GetOrderAllPendingOrderClient, error)
	}

	defaultOrderService struct {
		cli zrpc.Client
	}
)

func NewOrderService(cli zrpc.Client) OrderService {
	return &defaultOrderService{
		cli: cli,
	}
}

// 下单
func (m *defaultOrderService) Order(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*OrderEmpty, error) {
	client := pb.NewOrderServiceClient(m.cli.Conn())
	return client.Order(ctx, in, opts...)
}

// 创建订单,下单有分布式事务要处理分为两个接口
func (m *defaultOrderService) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*OrderEmpty, error) {
	client := pb.NewOrderServiceClient(m.cli.Conn())
	return client.CreateOrder(ctx, in, opts...)
}

// 获取用户订单列表
func (m *defaultOrderService) GetOrderList(ctx context.Context, in *GetOrderListByUserReq, opts ...grpc.CallOption) (*GetOrderListByUserResp, error) {
	client := pb.NewOrderServiceClient(m.cli.Conn())
	return client.GetOrderList(ctx, in, opts...)
}

// 取消订单
func (m *defaultOrderService) CancelOrder(ctx context.Context, in *CancelOrderReq, opts ...grpc.CallOption) (*OrderEmpty, error) {
	client := pb.NewOrderServiceClient(m.cli.Conn())
	return client.CancelOrder(ctx, in, opts...)
}

// 下单补偿
func (m *defaultOrderService) CreateOrderRevert(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*OrderEmpty, error) {
	client := pb.NewOrderServiceClient(m.cli.Conn())
	return client.CreateOrderRevert(ctx, in, opts...)
}

// 获取所有订单状态为未成交或部分成交的订单
func (m *defaultOrderService) GetOrderAllPendingOrder(ctx context.Context, in *OrderEmpty, opts ...grpc.CallOption) (pb.OrderService_GetOrderAllPendingOrderClient, error) {
	client := pb.NewOrderServiceClient(m.cli.Conn())
	return client.GetOrderAllPendingOrder(ctx, in, opts...)
}
