package itemHandler

import (
	"github.com/tonrock01/fantasia-shop/config"
	"github.com/tonrock01/fantasia-shop/modules/item/itemUsecase"
)

type (
	ItemHtppHandlerService interface{}

	itemHtppHandler struct {
		cfg         *config.Config
		itemUsecase itemUsecase.ItemUsecaseService
	}
)

func NewItemHttpHandler(cfg *config.Config, itemUsecase itemUsecase.ItemUsecaseService) ItemHtppHandlerService {
	return &itemHtppHandler{cfg, itemUsecase}
}
