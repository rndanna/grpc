package client

import (
	"fmt"

	"google.golang.org/grpc"
	"honnef.co/go/tools/config"
)

type carClient struct {
	client pb.ProductServiceClient
}

func NewCarClient(cfg config.Config) interfaces.ProductClient {
	fmt.Println("product client: ", cfg.ProductSvcUrl)
	grpcConnectoin, err := grpc.Dial(cfg.ProductSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewProductServiceClient(grpcConnectoin)

	return &productClient{
		client: grpcClient,
	}
}
