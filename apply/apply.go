package apply

import (
	"github.com/gofiber/fiber/v2"
	"kredit-plus/src/infrastructure/router"
	serviceBilling "kredit-plus/src/infrastructure/services/billings"
	serviceLoan "kredit-plus/src/infrastructure/services/loans"
	serviceTenor "kredit-plus/src/infrastructure/services/master_tenor"
	serviceUser "kredit-plus/src/infrastructure/services/users"
	container "kredit-plus/src/shared/di"
	"kredit-plus/src/usecase/billings"
	"kredit-plus/src/usecase/loans"
	"kredit-plus/src/usecase/master_tenor"
	"kredit-plus/src/usecase/users"
)

func Apply(f fiber.Router, ctn *container.Container) {
	router.NewUsers(f, serviceUser.NewUserService(ctn.Resolve("users-usecase").(*users.UsersUsecase)))
	router.NewBillings(f, serviceBilling.NewBillingsService(ctn.Resolve("billings-usecase").(*billings.BillingsUsecase)))
	router.Newloans(f, serviceLoan.NewLoansService(ctn.Resolve("loans-usecase").(*loans.LoansUsecase)))
	router.NewMasterTenor(f, serviceTenor.NewMasterTenorService(ctn.Resolve("tenor-usecase").(*master_tenor.MasterTenorUsecase)))
}
