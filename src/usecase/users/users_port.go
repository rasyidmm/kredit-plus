package users

import (
	"context"
	repoTenor "kredit-plus/src/adapter/repository/master_tenor"
	repoUser "kredit-plus/src/adapter/repository/users"
	"kredit-plus/src/usecase/payloads"
)

type UsersUsecase struct {
	repo      repoUser.UsersRepository
	repoTenor repoTenor.MasterTenorRepository
}

func NewUsersUsecase(r repoUser.UsersRepository, rt repoTenor.MasterTenorRepository) *UsersUsecase {
	return &UsersUsecase{repo: r, repoTenor: rt}

}

type UsersPort interface {
	Registration(ctx context.Context, user *payloads.RegistrationRequest) error
	GetUserList(ctx context.Context, req *payloads.GetUserListRequest) (*payloads.GetUserListResponse, error)
	Token(ctx context.Context, nik string) (*string, error)
}
