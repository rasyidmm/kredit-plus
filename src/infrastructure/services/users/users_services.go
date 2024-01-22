package users

import (
	"github.com/gofiber/fiber/v2"
	"kredit-plus/src/usecase/payloads"
	uce "kredit-plus/src/usecase/users"
	"strconv"
)

type UsersService struct {
	usecase uce.UsersPort
}

func NewUserService(u uce.UsersPort) *UsersService {
	return &UsersService{usecase: u}
}

func (s *UsersService) Registration(ctx *fiber.Ctx) error {
	var req payloads.RegistrationRequest

	if err := ctx.BodyParser(&req); nil != err {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	err := s.usecase.Registration(ctx.Context(), &req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON("success")
}

func (s *UsersService) GetUserList(ctx *fiber.Ctx) error {
	curPage, _ := strconv.Atoi(ctx.Query("cur_page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	req := payloads.GetUserListRequest{
		Status:  ctx.Query("status"),
		OrderBy: ctx.Query("order_by"),
		CurPage: curPage,
		Limit:   limit,
	}

	data, err := s.usecase.GetUserList(ctx.Context(), &req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(data)

}

func (s *UsersService) Token(ctx *fiber.Ctx) error {
	nik := ctx.Params("nik")

	data, err := s.usecase.Token(ctx.Context(), nik)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(data)

}
