package users

import (
	"context"
	"kredit-plus/src/adapter/model"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/utils"
)

type UsersRepository interface {
	Registration(ctx context.Context, user *entity.UsersEntity) error
	GetUserById(ctx context.Context, id int64) (*entity.UsersEntity, error)
	GetUserByNik(ctx context.Context, nik string) (*entity.UsersEntity, error)
	GetUserList(ctx context.Context, req *model.GetUserListRequest) (*[]entity.UsersEntity, *utils.PaginationHelper, error)
}
