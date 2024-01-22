package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"kredit-plus/src/adapter/model"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/utils"
)

type MockUsersRepository struct {
	mock.Mock
}

func (m *MockUsersRepository) Registration(ctx context.Context, user *entity.UsersEntity) error {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		return args.Error(1)
	}
	return nil
}
func (m *MockUsersRepository) GetUserById(ctx context.Context, id int64) (*entity.UsersEntity, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.UsersEntity), nil
}
func (m *MockUsersRepository) GetUserByNik(ctx context.Context, nik string) (*entity.UsersEntity, error) {
	args := m.Called(ctx, nik)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.UsersEntity), nil
}
func (m *MockUsersRepository) GetUserList(ctx context.Context, req *model.GetUserListRequest) (*[]entity.UsersEntity, *utils.PaginationHelper, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, nil, args.Error(1)
	}
	return args.Get(0).(*[]entity.UsersEntity), args.Get(1).(*utils.PaginationHelper), nil
}
