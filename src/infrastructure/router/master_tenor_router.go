package router

import (
	"github.com/gofiber/fiber/v2"
	"kredit-plus/src/infrastructure/services/master_tenor"
)

func NewMasterTenor(f fiber.Router, service *master_tenor.MasterTenorService) *fiber.Router {
	f.Group("tenor")
	return &f
}
