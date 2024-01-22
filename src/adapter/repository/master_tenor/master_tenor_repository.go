package master_tenor

import (
	"context"
	"kredit-plus/src/adapter/repository/entity"
)

type MasterTenorRepository interface {
	GetMasterTenorById(ctx context.Context, id int64) (*entity.MasterTenorEntity, error)
	GetMasterTenorByUserId(ctx context.Context, id int64) (*[]entity.MasterTenorEntity, error)
	GetMasterTenorByUserIdAndTenor(ctx context.Context, id int64, tenor int) (*entity.MasterTenorEntity, error)
}
