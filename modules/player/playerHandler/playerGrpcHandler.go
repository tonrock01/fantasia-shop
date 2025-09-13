package playerHandler

import (
	"github.com/tonrock01/fantasia-shop/modules/player/playerUsecase"
)

type (
	playerGrpcHandlerService struct {
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerGrpcHandler(playerUsecase playerUsecase.PlayerUsecaseService) *playerGrpcHandlerService {
	return &playerGrpcHandlerService{playerUsecase}
}
