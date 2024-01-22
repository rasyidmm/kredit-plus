package users

import (
	"context"
	"gorm.io/gorm"
	"kredit-plus/src/adapter/model"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/utils"
	"kredit-plus/src/shared/utils/constant"
	"strings"
)

type UsersDataHandler struct {
	db *gorm.DB
}

func NewUsersDataHandler(db *gorm.DB) *UsersDataHandler {
	return &UsersDataHandler{db: db}
}

func (d *UsersDataHandler) Registration(ctx context.Context, user *entity.UsersEntity) error {

	err := d.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *UsersDataHandler) GetUserById(ctx context.Context, id int64) (*entity.UsersEntity, error) {
	var user *entity.UsersEntity
	err := d.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *UsersDataHandler) GetUserByNik(ctx context.Context, nik string) (*entity.UsersEntity, error) {
	var user *entity.UsersEntity
	err := d.db.WithContext(ctx).Where("nik = ?", nik).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (d *UsersDataHandler) GetUserList(ctx context.Context, req *model.GetUserListRequest) (*[]entity.UsersEntity, *utils.PaginationHelper, error) {
	q := d.db.WithContext(ctx).Table(entity.UsersEntity{}.TableName())

	if req.Status != "" && constant.StatusUserOrderBy[strings.ToUpper(req.Status)] {
		q = q.Where("status = ?", strings.ToUpper(req.Status))
	}

	orderReq := strings.Split(req.OrderBy, ".")
	if len(orderReq) == 2 && (constant.UserIsOrderBy[strings.ToLower(orderReq[0])] && constant.OrderAction[strings.ToLower(orderReq[1])]) {
		q = q.Order(orderReq[0] + " " + orderReq[1])
	} else {
		q = q.Order("created_at desc")
	}

	var countAll int64
	q.Count(&countAll)

	pagination := utils.MakePagination(req.CurPage, req.Limit, int(countAll))
	q = q.Offset((req.CurPage - 1) * req.Limit)

	if req.Limit != 0 {
		q = q.Limit(pagination.PerPage)
	}

	var data *[]entity.UsersEntity
	err := q.Find(&data).Error
	if err != nil {
		return nil, nil, err
	}

	return data, pagination, nil
}
