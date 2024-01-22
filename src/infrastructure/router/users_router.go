package router

import (
	"github.com/gofiber/fiber/v2"
	"kredit-plus/src/infrastructure/services/users"
	"kredit-plus/src/shared/jwt"
)

func NewUsers(f fiber.Router, service *users.UsersService) *fiber.Router {
	r := f.Group("users")
	r.Post("/registrations", service.Registration)
	r.Get("", jwt.Authentication(service.GetUserList))
	r.Get("/:nik/token", service.Token)

	return &f

}
