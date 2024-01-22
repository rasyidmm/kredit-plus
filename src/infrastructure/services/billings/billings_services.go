package billings

import (
	"github.com/gofiber/fiber/v2"
	uce "kredit-plus/src/usecase/billings"
	"kredit-plus/src/usecase/payloads"
	"strconv"
)

type BillingsService struct {
	usecase uce.BillingsPort
}

func NewBillingsService(u uce.BillingsPort) *BillingsService {
	return &BillingsService{usecase: u}
}

func (s *BillingsService) BillingPayment(ctx *fiber.Ctx) error {

	var req payloads.BillingPaymentRequest
	if err := ctx.BodyParser(&req); nil != err {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := s.usecase.BillingPayment(ctx.Context(), &req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON("success")
}

func (s *BillingsService) GetBillingList(ctx *fiber.Ctx) error {

	curPage, _ := strconv.Atoi(ctx.Query("cur_page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	loadId, _ := strconv.Atoi(ctx.Query("loan_id"))

	req := payloads.BillingListRequest{
		Status:  ctx.Query("status"),
		OrderBy: ctx.Query("order_by"),
		LoanId:  int64(loadId),
		UserId:  int64(userId),
		CurPage: curPage,
		Limit:   limit,
	}

	data, err := s.usecase.GetBillingList(ctx.Context(), &req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(data)

}

func (s *BillingsService) GetBillingById(ctx *fiber.Ctx) error {
	loanId, _ := strconv.Atoi(ctx.Params("id"))
	data, err := s.usecase.GetBillingById(ctx.Context(), int64(loanId))
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(data)

}

func (s *BillingsService) BillingPaymentApprove(ctx *fiber.Ctx) error {
	loanId, _ := strconv.Atoi(ctx.Params("id"))
	err := s.usecase.BillingPaymentApprove(ctx.Context(), int64(loanId))
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON("success")
}

func (s *BillingsService) BillingPaymentReject(ctx *fiber.Ctx) error {
	loanId, _ := strconv.Atoi(ctx.Params("id"))
	err := s.usecase.BillingPaymentReject(ctx.Context(), int64(loanId))
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON("success")

}
