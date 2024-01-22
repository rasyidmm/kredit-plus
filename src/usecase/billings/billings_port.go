package billings

import (
	"context"
	repoBilling "kredit-plus/src/adapter/repository/billings"
	"kredit-plus/src/adapter/repository/entity"
	repoLoan "kredit-plus/src/adapter/repository/loans"
	"kredit-plus/src/usecase/payloads"
)

type BillingsUsecase struct {
	repo     repoBilling.BillingRepository
	repoLoan repoLoan.LoanRepository
}

func NewBillingsUsecase(r repoBilling.BillingRepository, rl repoLoan.LoanRepository) *BillingsUsecase {
	return &BillingsUsecase{repo: r, repoLoan: rl}

}

type BillingsPort interface {
	BillingPayment(ctx context.Context, req *payloads.BillingPaymentRequest) error
	GetBillingList(ctx context.Context, req *payloads.BillingListRequest) (*payloads.BillingListResponse, error)
	BillingPaymentApprove(ctx context.Context, id int64) error
	BillingPaymentReject(ctx context.Context, id int64) error
	GetBillingById(ctx context.Context, id int64) (*entity.BillingsEntity, error)
}
