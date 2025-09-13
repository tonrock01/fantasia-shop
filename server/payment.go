package server

import (
	"github.com/tonrock01/fantasia-shop/modules/payment/paymentHandler"
	"github.com/tonrock01/fantasia-shop/modules/payment/paymentRepository"
	"github.com/tonrock01/fantasia-shop/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	repo := paymentRepository.NewPaymentRepository(s.db)
	usecase := paymentUsecase.NewPaymentUsecase(repo)
	httpHandler := paymentHandler.NewPaymentHttpHandler(s.cfg, usecase)
	queueHandler := paymentHandler.NewPaymentQueueHandler(usecase)

	_ = httpHandler
	_ = queueHandler

	payment := s.app.Group("/payment_v1")

	// Health Check
	_ = payment
}
