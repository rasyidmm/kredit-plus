package entity

type MasterTenorEntity struct {
	Id       *int64  `gorm:"column:id"`
	UserId   int64   `gorm:"column:user_id"`
	Tenor    int64   `gorm:"column:tenor"`
	Amount   float64 `gorm:"column:amount"`
	Interest float64 `gorm:"column:interest"`
}

func (MasterTenorEntity) TableName() string {
	return "master_tenor"
}
