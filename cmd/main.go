package main

import (
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/golibbek/bookshop_order/config"
	pb "github.com/golibbek/bookshop_order/genproto/order_service"
	"github.com/golibbek/bookshop_order/pkg/db"
	"github.com/golibbek/bookshop_order/pkg/logger"
	"github.com/golibbek/bookshop_order/service"
	client "github.com/golibbek/bookshop_order/service/grpc_client"
	"github.com/golibbek/bookshop_order/storage"

)



func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "order-service")
	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {

		log.Fatal("sqlx connection to postgres error", logger.Error(err))

	}




	client, err := client.New(cfg)




	pgStorage := storage.NewStoragePg(connDB, client)




	orderService := service.NewOrderService(pgStorage, log)




	if err != nil {

		log.Error("gRPC dial error", logger.Error(err))

	}




	lis, err := net.Listen("tcp", cfg.RPCPort)

	if err != nil {

		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, orderService)
	reflection.Register(s)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}