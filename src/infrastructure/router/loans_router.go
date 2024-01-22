package router

import (
	"github.com/gofiber/fiber/v2"
	"kredit-plus/src/infrastructure/services/loans"
	"kredit-plus/src/shared/jwt"
)

func Newloans(f fiber.Router, service *loans.LoansService) *fiber.Router {
	r := f.Group("loans")
	r.Post("/submission", jwt.Authentication(service.SubmissionLoan))
	r.Get("", jwt.Authentication(service.GetLoanList))
	r.Get("/:id", jwt.Authentication(service.GetLoanById))
	r.Post("/approve/:id", jwt.Authentication(service.LoanApprove))
	r.Post("/reject/:id", jwt.Authentication(service.LoanReject))
	return &f
}
