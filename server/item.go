package server

import (
	"log"

	"github.com/tonrock01/fantasia-shop/modules/item/itemHandler"
	itemPb "github.com/tonrock01/fantasia-shop/modules/item/itemPb"
	"github.com/tonrock01/fantasia-shop/modules/item/itemRepository"
	"github.com/tonrock01/fantasia-shop/modules/item/itemUsecase"
	"github.com/tonrock01/fantasia-shop/pkg/grpcconn"
)

func (s *server) itemService() {
	repo := itemRepository.NewItemRepository(s.db)
	usecase := itemUsecase.NewItemUsecase(repo)
	httpHandler := itemHandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemHandler.NewItemGrpcHandler(usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpcconn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.ItemUrl)

		itemPb.RegisterItemGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Item gRPC server listening on %s", s.cfg.Grpc.ItemUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item_v1")

	// Health Check
	item.GET("", s.healthCheckService)
}
