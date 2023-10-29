package grpc

import (
	"context"
	gatewayService "gateway-service/proto/gateway"
)

type gatewayGRPCService struct {
	gatewayService.UnimplementedGatewayServiceServer
	carClient client.carClient
	engineClient ecgineClient
}

func New() *gatewayGRPCService {
	return &gatewayGRPCService{}
}

func (g *gatewayGRPCService) GetUserCars(ctx context.Context, in *gatewayService.GetUserCarsReq) (*gatewayService.GetUserCarsRes, error) {
	var res gatewayService.GetUserCarsRes

	return &res, nil
}

func (g *gatewayGRPCService) GetUserEngines(ctx context.Context, in *gatewayService.GetUserCarsReq) (*gatewayService.GetUserEnginesRes, error) {
	return nil, nil
}

func (g *gatewayGRPCService) GetUserEnginesByBrand(ctx context.Context, in *gatewayService.GetUserCarsByBrandReq) (*gatewayService.GetUserEnginesRes, error) {
	return nil, nil
}

func (g *gatewayGRPCService) GetUserEngineByCar(ctx context.Context, in *gatewayService.GetUserCarsReq) (*gatewayService.GetUserEngineRes, error) {
	return nil, nil
}
