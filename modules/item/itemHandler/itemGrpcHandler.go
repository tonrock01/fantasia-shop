package itemHandler

import (
	"github.com/tonrock01/fantasia-shop/modules/item/itemUsecase"
)

type (
	itemGrpcHandler struct {
		itemUsecase itemUsecase.ItemUsecaseService
	}
)

func NewItemGrpcHandler(itemUsecase itemUsecase.ItemUsecaseService) *itemGrpcHandler {
	return &itemGrpcHandler{itemUsecase}
}
