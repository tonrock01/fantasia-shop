package server

import (
	"log"

	"github.com/tonrock01/fantasia-shop/modules/player/playerHandler"
	playerPb "github.com/tonrock01/fantasia-shop/modules/player/playerPb"
	"github.com/tonrock01/fantasia-shop/modules/player/playerRepository"
	"github.com/tonrock01/fantasia-shop/modules/player/playerUsecase"
	"github.com/tonrock01/fantasia-shop/pkg/grpcconn"
)

func (s *server) playerService() {
	repo := playerRepository.NewPlayerRepository(s.db)
	usecase := playerUsecase.NewPlayerUsecase(repo)
	httpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, usecase)
	grpcHandler := playerHandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerHandler.NewPlayeQueuerHandler(s.cfg, usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpcconn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.PlayerUrl)

		playerPb.RegisterPlayerGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Player gRPC server listening on %s", s.cfg.Grpc.PlayerUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/player_v1")

	// Health Check
	player.GET("", s.healthCheckService)
}
