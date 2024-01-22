package entity

import "time"

type UsersEntity struct {
	Id                  *int64    `gorm:"column:id"`
	Nik                 string    `gorm:"column:nik"`
	FullName            string    `gorm:"column:full_name"`
	LegalName           string    `gorm:"column:legal_name"`
	Birthplace          string    `gorm:"column:birthplace"`
	Birthdate           time.Time `gorm:"column:birthdate"`
	Salary              float64   `gorm:"column:salary"`
	IdentificationPhoto string    `gorm:"column:identification_photo"`
	PhotoSelfie         string    `gorm:"column:photo_selfie"`
	Status              string    `gorm:"column:status"`
	CreatedAt           time.Time `gorm:"column:created_at"`
}

func (UsersEntity) TableName() string {
	return "users"
}
