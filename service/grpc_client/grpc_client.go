package grpcClient

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/golibbek/bookshop_order/config"
	pb "github.com/golibbek/bookshop_order/genproto/catalog_service"
)

// IServiceManager ...
type IServiceManager interface {
	CatalogService() pb.CatalogServiceClient
}

type serviceManager struct {
	catalogService pb.CatalogServiceClient
}

func (s *serviceManager) CatalogService() pb.CatalogServiceClient {
	return s.catalogService
}

// New ...
func New(cfg config.Config) (IServiceManager, error) {
	connCatalog, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CatalogServiceHost, cfg.CatalogServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		catalogService: pb.NewCatalogServiceClient(connCatalog),
	}

	return serviceManager, nil
}