package loans

import (
	"context"
	"kredit-plus/src/adapter/repository/entity"
	repoLoans "kredit-plus/src/adapter/repository/loans"
	repoTenor "kredit-plus/src/adapter/repository/master_tenor"
	repoUser "kredit-plus/src/adapter/repository/users"
	"kredit-plus/src/usecase/payloads"
)

type LoansUsecase struct {
	repo      repoLoans.LoanRepository
	repoTenor repoTenor.MasterTenorRepository
	repoUser  repoUser.UsersRepository
}

func NewLoansUsecase(r repoLoans.LoanRepository, rt repoTenor.MasterTenorRepository, ru repoUser.UsersRepository) *LoansUsecase {
	return &LoansUsecase{repo: r, repoTenor: rt, repoUser: ru}

}

type LoansPort interface {
	SubmissionLoans(ctx context.Context, req *payloads.SubmissionLoanRequest) error
	GetLoanList(ctx context.Context, req *payloads.LoanListRequest) (*payloads.GetLoanListResponse, error)
	GetLoanById(ctx context.Context, id int64) (*entity.LoansEntity, error)
	LoanApprove(ctx context.Context, id int64) error
	LoanReject(ctx context.Context, id int64) error
}
