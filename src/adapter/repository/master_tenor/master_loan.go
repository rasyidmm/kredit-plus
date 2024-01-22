package master_tenor

import (
	"context"
	"gorm.io/gorm"
	"kredit-plus/src/adapter/repository/entity"
)

type MasterTenorDataHandler struct {
	db *gorm.DB
}

func NewMasterTenorDataHandler(db *gorm.DB) *MasterTenorDataHandler {
	return &MasterTenorDataHandler{db: db}
}

func (d *MasterTenorDataHandler) GetMasterTenorById(ctx context.Context, id int64) (*entity.MasterTenorEntity, error) {
	var masterTenor *entity.MasterTenorEntity
	err := d.db.WithContext(ctx).Where("id = ?", id).First(&masterTenor).Error
	if err != nil {
		return nil, err
	}
	return masterTenor, nil

}
func (d *MasterTenorDataHandler) GetMasterTenorByUserId(ctx context.Context, id int64) (*[]entity.MasterTenorEntity, error) {
	var masterTenor *[]entity.MasterTenorEntity
	err := d.db.WithContext(ctx).Where("user_id = ?", id).Find(&masterTenor).Error
	if err != nil {
		return nil, err
	}
	return masterTenor, nil

}

func (d *MasterTenorDataHandler) GetMasterTenorByUserIdAndTenor(ctx context.Context, id int64, tenor int) (*entity.MasterTenorEntity, error) {
	var masterTenor *entity.MasterTenorEntity
	err := d.db.WithContext(ctx).Where("user_id = ? and tenor = ?", id, tenor).First(&masterTenor).Error
	if err != nil {
		return nil, err
	}
	return masterTenor, nil

}
