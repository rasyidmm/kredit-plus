package loans

import (
	"github.com/gofiber/fiber/v2"
	uce "kredit-plus/src/usecase/loans"
	"kredit-plus/src/usecase/payloads"
	"strconv"
)

type LoansService struct {
	usecase uce.LoansPort
}

func NewLoansService(u uce.LoansPort) *LoansService {
	return &LoansService{usecase: u}
}

func (s *LoansService) SubmissionLoan(ctx *fiber.Ctx) error {
	var req payloads.SubmissionLoanRequest

	if err := ctx.BodyParser(&req); nil != err {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := s.usecase.SubmissionLoans(ctx.Context(), &req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON("success")
}

func (s *LoansService) GetLoanList(ctx *fiber.Ctx) error {
	curPage, _ := strconv.Atoi(ctx.Query("cur_page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	userId, _ := strconv.Atoi(ctx.Query("user_id"))

	req := payloads.LoanListRequest{
		Status:  ctx.Query("status"),
		OrderBy: ctx.Query("order_by"),
		UserId:  int64(userId),
		CurPage: curPage,
		Limit:   limit,
	}

	data, err := s.usecase.GetLoanList(ctx.Context(), &req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(fiber.StatusOK).JSON(data)

}

func (s *LoansService) GetLoanById(ctx *fiber.Ctx) error {
	loanId, _ := strconv.Atoi(ctx.Params("id"))
	data, err := s.usecase.GetLoanById(ctx.Context(), int64(loanId))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(fiber.StatusOK).JSON(data)

}

func (s *LoansService) LoanApprove(ctx *fiber.Ctx) error {
	loanId, _ := strconv.Atoi(ctx.Params("id"))
	err := s.usecase.LoanApprove(ctx.Context(), int64(loanId))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(fiber.StatusOK).JSON("success")

}

func (s *LoansService) LoanReject(ctx *fiber.Ctx) error {
	loanId, _ := strconv.Atoi(ctx.Params("id"))
	err := s.usecase.LoanReject(ctx.Context(), int64(loanId))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(fiber.StatusOK).JSON("success")

}
