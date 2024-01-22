package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"kredit-plus/src/adapter/model"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/utils"
)

type MockBillingRepository struct {
	mock.Mock
}

func (m *MockBillingRepository) CreateBilling(ctx context.Context, req *entity.BillingsEntity) error {
	args := m.Called(ctx, req)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (m *MockBillingRepository) GetBillingListByLoanIdAndStatus(ctx context.Context, loanId int64, status []string) (*[]entity.BillingsEntity, error) {

	args := m.Called(ctx, loanId, status)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]entity.BillingsEntity), nil
}
func (m *MockBillingRepository) GetBillingList(ctx context.Context, req *model.BillingListRequest) (*[]entity.BillingsEntity, *utils.PaginationHelper, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, nil, args.Error(2)
	}
	return args.Get(0).(*[]entity.BillingsEntity), args.Get(1).(*utils.PaginationHelper), nil
}
func (m *MockBillingRepository) GetBillingById(ctx context.Context, id int64) (*entity.BillingsEntity, error) {

	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.BillingsEntity), nil
}
func (m *MockBillingRepository) UpdateBilling(ctx context.Context, bill *entity.BillingsEntity) error {

	args := m.Called(ctx, bill)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
