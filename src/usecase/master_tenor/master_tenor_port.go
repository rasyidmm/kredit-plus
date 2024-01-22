package master_tenor

import (
	repoMasterTenor "kredit-plus/src/adapter/repository/master_tenor"
)

type MasterTenorUsecase struct {
	repo repoMasterTenor.MasterTenorRepository
}

func NewMasterTenorUsecase(repo repoMasterTenor.MasterTenorRepository) *MasterTenorUsecase {
	return &MasterTenorUsecase{repo: repo}

}

type MasterTenorPort interface {
}
