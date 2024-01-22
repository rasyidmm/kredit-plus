package loans

import (
	"context"
	"gorm.io/gorm"
	"kredit-plus/src/adapter/model"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/utils"
	"kredit-plus/src/shared/utils/constant"
	"strings"
)

type LoansDataHandler struct {
	db *gorm.DB
}

func NewLoansDataHandler(db *gorm.DB) *LoansDataHandler {
	return &LoansDataHandler{db: db}
}

func (d *LoansDataHandler) GetLoanListByUserIdAndTenor(ctx context.Context, id int64, tenor int, status []string) (*[]entity.LoansEntity, error) {
	var loans *[]entity.LoansEntity
	err := d.db.WithContext(ctx).Where("user_id = ? and tenor = ? and status in (?)", id, tenor, status).Find(&loans).Error
	if err != nil {
		return nil, err
	}
	return loans, nil
}

func (d *LoansDataHandler) CreateLoan(ctx context.Context, loan *entity.LoansEntity) error {
	err := d.db.WithContext(ctx).Create(loan).Error
	if err != nil {
		return err
	}
	return nil

}

func (d *LoansDataHandler) GetLoanList(ctx context.Context, req *model.LoanListRequest) (*[]entity.LoansEntity, *utils.PaginationHelper, error) {
	q := d.db.WithContext(ctx).Table(entity.LoansEntity{}.TableName())

	if req.Status != "" && constant.StatusLoanOrderBy[strings.ToUpper(req.Status)] {
		q = q.Where("status = ?", strings.ToUpper(req.Status))
	}

	if req.UserId != 0 {
		q = q.Where("user_id = ?", req.UserId)
	}

	orderReq := strings.Split(req.OrderBy, ".")
	if len(orderReq) == 2 && (constant.LoanIsOrderBy[strings.ToLower(orderReq[0])] && constant.OrderAction[strings.ToLower(orderReq[1])]) {
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

	var data *[]entity.LoansEntity
	err := q.Find(&data).Error
	if err != nil {
		return nil, nil, err
	}

	return data, pagination, nil
}

func (d *LoansDataHandler) GetLoanById(ctx context.Context, id int64) (*entity.LoansEntity, error) {
	var loan *entity.LoansEntity
	err := d.db.WithContext(ctx).Where("id = ?", id).Find(&loan).Error
	if err != nil {
		return nil, err
	}
	return loan, nil
}

func (d *LoansDataHandler) UpdateLoan(ctx context.Context, loan *entity.LoansEntity) error {
	err := d.db.WithContext(ctx).Model(loan).Updates(loan).Error
	if err != nil {
		return err
	}
	return nil

}
