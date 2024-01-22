package entity

import "time"

type LoansEntity struct {
	Id                int64      `gorm:"column:id"`
	UserId            int64      `gorm:"column:user_id"`
	LoanNo            string     `gorm:"column:loan_no"`
	Otr               float64    `gorm:"column:otr"`
	AdminFee          float64    `gorm:"column:admin_fee"`
	InstallmentAmount float64    `gorm:"column:installment_amount"`
	InterestAmount    float64    `gorm:"column:interest_amount"`
	AssetName         string     `gorm:"column:asset_name"`
	Status            string     `gorm:"column:status"`
	Tenor             int        `gorm:"column:tenor"`
	CreatedAt         time.Time  `gorm:"column:created_at"`
	ApprovedAt        *time.Time `gorm:"column:approved_at"`
}

func (LoansEntity) TableName() string {
	return "loans"
}
