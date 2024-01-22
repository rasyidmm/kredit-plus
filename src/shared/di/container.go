package di

import (
	"github.com/sarulabs/di/v2"
	"kredit-plus/src/adapter/db"
	repositoryBillings "kredit-plus/src/adapter/repository/billings"
	repositoryloans "kredit-plus/src/adapter/repository/loans"
	repositoryTenor "kredit-plus/src/adapter/repository/master_tenor"
	repositoryUser "kredit-plus/src/adapter/repository/users"
	"kredit-plus/src/usecase/billings"
	"kredit-plus/src/usecase/loans"
	"kredit-plus/src/usecase/master_tenor"
	"kredit-plus/src/usecase/users"
)

type Container struct {
	ctn di.Container
}

func NewContainer() *Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add([]di.Def{
		{Name: "users-usecase", Build: usersUsecase},
		{Name: "billings-usecase", Build: billingsUsecase},
		{Name: "loans-usecase", Build: loansUsecase},
		{Name: "tenor-usecase", Build: masterTenorUsecase},
	}...)
	return &Container{
		ctn: builder.Build(),
	}
}
func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func usersUsecase(_ di.Container) (interface{}, error) {
	repoUser := repositoryUser.NewUsersDataHandler(db.DB)
	repoTenor := repositoryTenor.NewMasterTenorDataHandler(db.DB)
	return users.NewUsersUsecase(repoUser, repoTenor), nil

}
func billingsUsecase(_ di.Container) (interface{}, error) {
	repo := repositoryBillings.NewBillingsDataHandler(db.DB)
	repoLoan := repositoryloans.NewLoansDataHandler(db.DB)
	return billings.NewBillingsUsecase(repo, repoLoan), nil

}
func loansUsecase(_ di.Container) (interface{}, error) {
	repo := repositoryloans.NewLoansDataHandler(db.DB)
	repoUser := repositoryUser.NewUsersDataHandler(db.DB)
	repoTenor := repositoryTenor.NewMasterTenorDataHandler(db.DB)
	return loans.NewLoansUsecase(repo, repoTenor, repoUser), nil

}
func masterTenorUsecase(_ di.Container) (interface{}, error) {
	repo := repositoryTenor.NewMasterTenorDataHandler(db.DB)
	return master_tenor.NewMasterTenorUsecase(repo), nil

}
