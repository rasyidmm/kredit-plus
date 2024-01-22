package loans

import (
	"context"
	"kredit-plus/src/adapter/model"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/utils"
)

type LoanRepository interface {
	GetLoanListByUserIdAndTenor(ctx context.Context, id int64, tenor int, status []string) (*[]entity.LoansEntity, error)
	CreateLoan(ctx context.Context, loan *entity.LoansEntity) error
	GetLoanList(ctx context.Context, req *model.LoanListRequest) (*[]entity.LoansEntity, *utils.PaginationHelper, error)
	GetLoanById(ctx context.Context, id int64) (*entity.LoansEntity, error)
	UpdateLoan(ctx context.Context, loan *entity.LoansEntity) error
}
