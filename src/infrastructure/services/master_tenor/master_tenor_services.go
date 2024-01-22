package master_tenor

import uce "kredit-plus/src/usecase/master_tenor"

type MasterTenorService struct {
	usecase uce.MasterTenorPort
}

func NewMasterTenorService(u uce.MasterTenorPort) *MasterTenorService {
	return &MasterTenorService{usecase: u}
}
