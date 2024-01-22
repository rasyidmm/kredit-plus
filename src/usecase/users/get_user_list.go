package users

import (
	"context"
	"kredit-plus/src/adapter/model"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/usecase/payloads"
)

func (u *UsersUsecase) GetUserList(ctx context.Context, req *payloads.GetUserListRequest) (*payloads.GetUserListResponse, error) {
	data, pagination, err := u.repo.GetUserList(ctx, &model.GetUserListRequest{
		Status:  req.Status,
		OrderBy: req.OrderBy,
		CurPage: req.CurPage,
		Limit:   req.Limit,
	})

	if err != nil {
		return nil, err
	}

	var result []payloads.UserResponse
	for _, item := range *data {
		dataTenor, err := u.repoTenor.GetMasterTenorByUserId(ctx, *item.Id)
		if err != nil {
			return nil, err
		}
		result = append(result, buildResponseUser(item, dataTenor))
	}

	return &payloads.GetUserListResponse{
		Data:             result,
		PaginationHelper: pagination,
	}, nil

}

func buildResponseUser(user entity.UsersEntity, tenors *[]entity.MasterTenorEntity) payloads.UserResponse {
	var dataTenor []payloads.MasterTenor
	for _, tenor := range *tenors {
		dataTenor = append(dataTenor, payloads.MasterTenor{
			Id:       tenor.Id,
			UserId:   tenor.UserId,
			Tenor:    tenor.Tenor,
			Amount:   tenor.Amount,
			Interest: tenor.Interest,
		})
	}

	return payloads.UserResponse{
		Id:                  user.Id,
		Nik:                 user.Nik,
		FullName:            user.FullName,
		LegalName:           user.LegalName,
		Birthplace:          user.Birthplace,
		Birthdate:           user.Birthdate,
		Salary:              user.Salary,
		IdentificationPhoto: user.IdentificationPhoto,
		PhotoSelfie:         user.PhotoSelfie,
		Status:              user.Status,
		CreatedAt:           user.CreatedAt,
		Tenor:               &dataTenor,
	}
}
