package billings

import (
	"context"
	"gorm.io/gorm"
	"kredit-plus/src/adapter/model"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/utils"
	"kredit-plus/src/shared/utils/constant"
	"strings"
)

type BillingsDataHandler struct {
	db *gorm.DB
}

func NewBillingsDataHandler(db *gorm.DB) *BillingsDataHandler {
	return &BillingsDataHandler{db: db}
}

func (d *BillingsDataHandler) CreateBilling(ctx context.Context, bill *entity.BillingsEntity) error {
	err := d.db.WithContext(ctx).Create(bill).Error
	if err != nil {
		return err
	}
	return nil

}

func (d *BillingsDataHandler) GetBillingListByLoanIdAndStatus(ctx context.Context, loanId int64, status []string) (*[]entity.BillingsEntity, error) {
	var billings *[]entity.BillingsEntity
	err := d.db.WithContext(ctx).Where("loan_id = ? AND status IN ?", loanId, status).Find(&billings).Error
	if err != nil {
		return nil, err
	}
	return billings, nil

}

func (d *BillingsDataHandler) GetBillingList(ctx context.Context, req *model.BillingListRequest) (*[]entity.BillingsEntity, *utils.PaginationHelper, error) {
	q := d.db.WithContext(ctx).Table(entity.BillingsEntity{}.TableName())

	if req.Status != "" && constant.StatusUserOrderBy[strings.ToUpper(req.Status)] {
		q = q.Where("status = ?", strings.ToUpper(req.Status))
	}

	if req.UserId != 0 {
		q = q.Where("user_id = ?", req.UserId)
	}

	if req.LoanId != 0 {
		q = q.Where("loan_id = ?", req.LoanId)
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

	var data *[]entity.BillingsEntity
	err := q.Find(&data).Error
	if err != nil {
		return nil, nil, err
	}

	return data, pagination, nil
}

func (d *BillingsDataHandler) GetBillingById(ctx context.Context, id int64) (*entity.BillingsEntity, error) {
	var bill *entity.BillingsEntity
	err := d.db.WithContext(ctx).Where("id = ?", id).Find(&bill).Error
	if err != nil {
		return nil, err
	}
	return bill, nil
}

func (d *BillingsDataHandler) UpdateBilling(ctx context.Context, bill *entity.BillingsEntity) error {
	err := d.db.WithContext(ctx).Model(bill).Updates(bill).Error
	if err != nil {
		return err
	}
	return nil

}
