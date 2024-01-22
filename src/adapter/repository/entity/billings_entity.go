package entity

import "time"

type BillingsEntity struct {
	Id         int64      `gorm:"column:id"`
	LoanId     int64      `gorm:"column:loan_id"`
	UserId     int64      `gorm:"column:user_id"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	Amount     float64    `gorm:"column:amount"`
	Status     string     `gorm:"column:status"`
	ApprovedAt *time.Time `gorm:"column:approved_at"`
}

func (BillingsEntity) TableName() string {
	return "billings"
}
