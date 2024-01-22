package billings

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"kredit-plus/src/adapter/repository/entity"
)

func (u *BillingsUsecase) GetBillingById(ctx context.Context, id int64) (*entity.BillingsEntity, error) {

	data, err := u.repo.GetBillingById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return data, nil
}
