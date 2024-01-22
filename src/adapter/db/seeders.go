package db

import (
	"errors"
	"gorm.io/gorm"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/utils/constant"
	"time"
)

func Seed(db *gorm.DB) error {
	err := db.First(&entity.UsersEntity{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		users := []*entity.UsersEntity{
			{
				Id:                  nil,
				Nik:                 "1234567890",
				FullName:            "Budi",
				LegalName:           "Budi",
				Birthplace:          "Jakarta",
				Birthdate:           time.Now(),
				Salary:              5000000,
				IdentificationPhoto: "",
				PhotoSelfie:         "",
				Status:              constant.Approved,
				CreatedAt:           time.Now(),
			},
			{
				Id:                  nil,
				Nik:                 "1234567891",
				FullName:            "Annisa",
				LegalName:           "Annisa",
				Birthplace:          "Jakarta",
				Birthdate:           time.Now(),
				Salary:              5000000,
				IdentificationPhoto: "",
				PhotoSelfie:         "",
				Status:              constant.Approved,
				CreatedAt:           time.Now().Add(time.Hour * 1),
			},
		}
		db.Create(users)
	}

	err = db.First(&entity.MasterTenorEntity{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tenor := []*entity.MasterTenorEntity{
			{
				UserId:   1,
				Tenor:    1,
				Amount:   100000,
				Interest: 0,
			},
			{
				UserId:   1,
				Tenor:    2,
				Amount:   200000,
				Interest: 2,
			},
			{
				UserId:   1,
				Tenor:    3,
				Amount:   500000,
				Interest: 3,
			},
			{
				UserId:   1,
				Tenor:    4,
				Amount:   70000,
				Interest: 4,
			},
			{
				UserId:   2,
				Tenor:    1,
				Amount:   1000000,
				Interest: 0,
			},
			{
				UserId:   2,
				Tenor:    2,
				Amount:   1200000,
				Interest: 2,
			},
			{
				UserId:   2,
				Tenor:    3,
				Amount:   1000000,
				Interest: 3,
			},
			{
				UserId:   2,
				Tenor:    4,
				Amount:   1000000,
				Interest: 4,
			},
		}
		db.Create(tenor)
	}
	return nil
}
