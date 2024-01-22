package loans

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"kredit-plus/src/shared/utils/constant"
	"time"
)

func (u *LoansUsecase) LoanApprove(ctx context.Context, id int64) error {
	data, err := u.repo.GetLoanById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	t := time.Now()
	data.Status = constant.Approved
	data.ApprovedAt = &t
	err = u.repo.UpdateLoan(ctx, data)
	if err != nil {
		return err
	}
	return nil

}
