package billings

import (
	"context"
	"kredit-plus/src/adapter/model"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/utils"
)

type BillingRepository interface {
	CreateBilling(ctx context.Context, req *entity.BillingsEntity) error
	GetBillingListByLoanIdAndStatus(ctx context.Context, loanId int64, status []string) (*[]entity.BillingsEntity, error)
	GetBillingList(ctx context.Context, req *model.BillingListRequest) (*[]entity.BillingsEntity, *utils.PaginationHelper, error)
	GetBillingById(ctx context.Context, id int64) (*entity.BillingsEntity, error)
	UpdateBilling(ctx context.Context, bill *entity.BillingsEntity) error
}
