package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"kredit-plus/src/adapter/model"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/utils"
)

type MockLoansRepository struct {
	mock.Mock
}

func (m *MockLoansRepository) GetLoanListByUserIdAndTenor(ctx context.Context, id int64, tenor int, status []string) (*[]entity.LoansEntity, error) {
	args := m.Called(ctx, id, tenor, status)
	if args.Get(0) == nil {
		return nil, args.Error(0)
	}
	return args.Get(0).(*[]entity.LoansEntity), nil
}
func (m *MockLoansRepository) CreateLoan(ctx context.Context, loan *entity.LoansEntity) error {
	args := m.Called(ctx, loan)
	if args == nil {
		return args.Error(0)
	}
	return nil
}
func (m *MockLoansRepository) GetLoanList(ctx context.Context, req *model.LoanListRequest) (*[]entity.LoansEntity, *utils.PaginationHelper, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, nil, args.Error(2)
	}
	return args.Get(0).(*[]entity.LoansEntity), args.Get(1).(*utils.PaginationHelper), nil
}
func (m *MockLoansRepository) GetLoanById(ctx context.Context, id int64) (*entity.LoansEntity, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.LoansEntity), nil
}
func (m *MockLoansRepository) UpdateLoan(ctx context.Context, loan *entity.LoansEntity) error {
	args := m.Called(ctx, loan)
	if args == nil {
		return args.Error(0)
	}
	return nil
}
