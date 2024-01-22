package loans

import (
	"context"
	"kredit-plus/src/adapter/repository/entity"
)

func (u *LoansUsecase) GetLoanById(ctx context.Context, id int64) (*entity.LoansEntity, error) {
	loan, err := u.repo.GetLoanById(ctx, id)
	if err != nil {
		return nil, err
	}
	return loan, nil
}
