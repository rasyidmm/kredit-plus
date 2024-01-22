package router

import (
	"github.com/gofiber/fiber/v2"
	"kredit-plus/src/infrastructure/services/billings"
	"kredit-plus/src/shared/jwt"
)

func NewBillings(f fiber.Router, service *billings.BillingsService) *fiber.Router {
	r := f.Group("billings")
	r.Get("", jwt.Authentication(service.GetBillingList))
	r.Get("/:id", jwt.Authentication(service.GetBillingById))
	r.Post("/reject/:id", jwt.Authentication(service.BillingPaymentReject))
	r.Post("/approve/:id", jwt.Authentication(service.BillingPaymentApprove))
	r.Post("/payment", jwt.Authentication(service.BillingPayment))
	return &f
}
