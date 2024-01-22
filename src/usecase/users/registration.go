package users

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/utils"
	"kredit-plus/src/usecase/payloads"
	"net/http"
	"strings"
	"time"
)

func (u *UsersUsecase) Registration(ctx context.Context, user *payloads.RegistrationRequest) error {

	userExit, err := u.repo.GetUserByNik(ctx, user.Nik)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	if userExit != nil {
		return fiber.NewError(http.StatusNotFound, "nik already exist")
	}

	birthReq := strings.Split(user.Birthdate, "-")
	dateString := utils.ValueOrZero(&birthReq[2]) + "-" + utils.ValueOrZero(&birthReq[1]) + "-" + utils.ValueOrZero(&birthReq[0])
	date, _ := time.Parse("2006-01-02", dateString)

	request := &entity.UsersEntity{
		Nik:                 user.Nik,
		FullName:            user.FullName,
		LegalName:           user.LegalName,
		Birthplace:          user.Birthplace,
		Birthdate:           date,
		Salary:              user.Salary,
		IdentificationPhoto: user.IdentificationPhoto,
		PhotoSelfie:         user.PhotoSelfie,
	}
	err = u.repo.Registration(ctx, request)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
