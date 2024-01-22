package billings

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"kredit-plus/src/shared/utils/constant"
)

func (u *BillingsUsecase) BillingPaymentReject(ctx context.Context, id int64) error {

	data, err := u.repo.GetBillingById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	data.Status = constant.Rejected
	err = u.repo.UpdateBilling(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
