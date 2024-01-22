package loans

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/utils/constant"
	"kredit-plus/src/usecase/payloads"
	"net/http"
	"time"
)

func (u *LoansUsecase) SubmissionLoans(ctx context.Context, req *payloads.SubmissionLoanRequest) error {
	user, err := u.repoUser.GetUserById(ctx, req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(http.StatusNotFound, "user Tidak Ditemukan")
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	tenor, err := u.repoTenor.GetMasterTenorByUserIdAndTenor(ctx, req.UserId, req.Tenor)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	loans, err := u.repo.GetLoanListByUserIdAndTenor(ctx, req.UserId, req.Tenor, []string{constant.InReview, constant.Approved})
	if err != nil {
		return err
	}
	var totalLoan float64
	if loans != nil {
		for _, itemLoan := range *loans {
			totalLoan += itemLoan.Otr - itemLoan.AdminFee
		}
	}

	if req.Amount+totalLoan > tenor.Amount {
		return fiber.NewError(http.StatusPreconditionFailed, "harga Melebihi Batas Tenor")
	}

	err = u.repo.CreateLoan(ctx, &entity.LoansEntity{
		UserId:            req.UserId,
		LoanNo:            user.Nik + "_" + time.Now().Format("20060102") + "_" + time.Now().Format("150405"),
		Otr:               (req.Amount + (req.Amount * tenor.Interest / 100)) + req.AdminFee,
		AdminFee:          req.AdminFee,
		InstallmentAmount: (req.Amount + (req.Amount * tenor.Interest / 100)) / float64(req.Tenor),
		InterestAmount:    tenor.Interest,
		AssetName:         req.AssetName,
		Status:            constant.InReview,
		Tenor:             req.Tenor,
		CreatedAt:         time.Now(),
	})

	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
