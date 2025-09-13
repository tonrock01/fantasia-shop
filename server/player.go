package server

import (
	"github.com/tonrock01/fantasia-shop/modules/player/playerHandler"
	"github.com/tonrock01/fantasia-shop/modules/player/playerRepository"
	"github.com/tonrock01/fantasia-shop/modules/player/playerUsecase"
)

func (s *server) playerService() {
	repo := playerRepository.NewPlayerRepository(s.db)
	usecase := playerUsecase.NewPlayerUsecase(repo)
	httpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, usecase)
	grpcHandler := playerHandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerHandler.NewPlayeQueuerHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/player_v1")

	// Health Check
	player.GET("", s.healthCheckService)
}
