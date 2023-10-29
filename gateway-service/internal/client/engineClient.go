package client


type engineClient struct {
	client pb.ProductServiceClient
}

func NewEngineClient(cfg config.Config) interfaces.ProductClient {
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
