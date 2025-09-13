package paymentHandler

import "github.com/tonrock01/fantasia-shop/modules/payment/paymentUsecase"

type (
	PaymentQueueHandlerService interface{}

	paymentQueueHandler struct {
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentQueueHandler(paymentUsecase paymentUsecase.PaymentUsecaseService) *paymentQueueHandler {
	return &paymentQueueHandler{paymentUsecase}
}
