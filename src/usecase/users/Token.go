package users

import (
	"context"
	"kredit-plus/src/shared/jwt"
)

func (u *UsersUsecase) Token(ctx context.Context, nik string) (*string, error) {
	user, err := u.repo.GetUserByNik(ctx, nik)
	if err != nil {
		return nil, err
	}

	token, err := jwt.GenerateToken(*user.Id)
	if err != nil {
		return nil, err
	}
	return &token, nil

}
