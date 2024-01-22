package billings

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/utils/constant"
	"kredit-plus/src/usecase/payloads"
	"net/http"
	"time"
)

func (u *BillingsUsecase) BillingPayment(ctx context.Context, req *payloads.BillingPaymentRequest) error {
	loan, err := u.repoLoan.GetLoanById(ctx, req.LoanId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	billings, err := u.repo.GetBillingListByLoanIdAndStatus(ctx, loan.Id, []string{constant.InReview, constant.Approved})
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	var totalAmount float64
	if len(*billings) > 0 {
		totalAmount = loan.InstallmentAmount
	} else {
		totalAmount = loan.InstallmentAmount + loan.AdminFee
	}

	if totalAmount != req.Amount {
		return fiber.NewError(http.StatusBadRequest, "total amount not match installment")
	}

	err = u.repo.CreateBilling(ctx, &entity.BillingsEntity{
		LoanId:    req.LoanId,
		UserId:    req.UserId,
		CreatedAt: time.Now(),
		Amount:    req.Amount,
		Status:    constant.InReview,
	})

	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
