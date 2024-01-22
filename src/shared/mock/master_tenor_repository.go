package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"kredit-plus/src/adapter/repository/entity"
)

type MockMasterTenorRepository struct {
	mock.Mock
}

func (m *MockMasterTenorRepository) GetMasterTenorById(ctx context.Context, id int64) (*entity.MasterTenorEntity, error) {
	argh := m.Called(ctx, id)
	if argh.Get(0) == nil {
		return nil, argh.Error(0)
	}
	return argh.Get(0).(*entity.MasterTenorEntity), nil
}
func (m *MockMasterTenorRepository) GetMasterTenorByUserId(ctx context.Context, id int64) (*[]entity.MasterTenorEntity, error) {
	argh := m.Called(ctx, id)
	if argh.Get(0) == nil {
		return nil, argh.Error(0)
	}
	return argh.Get(0).(*[]entity.MasterTenorEntity), nil
}
func (m *MockMasterTenorRepository) GetMasterTenorByUserIdAndTenor(ctx context.Context, id int64, tenor int) (*entity.MasterTenorEntity, error) {
	argh := m.Called(ctx, id, tenor)
	if argh.Get(0) == nil {
		return nil, argh.Error(0)
	}
	return argh.Get(0).(*entity.MasterTenorEntity), nil
}
