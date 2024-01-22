package loans

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"kredit-plus/src/adapter/model"
	"kredit-plus/src/usecase/payloads"
	"net/http"
)

func (u *LoansUsecase) GetLoanList(ctx context.Context, req *payloads.LoanListRequest) (*payloads.GetLoanListResponse, error) {
	data, pagination, err := u.repo.GetLoanList(ctx, &model.LoanListRequest{
		UserId:  req.UserId,
		Status:  req.Status,
		OrderBy: req.OrderBy,
		CurPage: req.CurPage,
		Limit:   req.Limit,
	})

	if err != nil {
		return nil, fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	var result []payloads.LoanResponse

	for _, item := range *data {
		result = append(result, payloads.LoanResponse{
			Id:                item.Id,
			UserId:            item.UserId,
			LoanNo:            item.LoanNo,
			Otr:               item.Otr,
			AdminFee:          item.AdminFee,
			InstallmentAmount: item.InstallmentAmount,
			InterestAmount:    item.InterestAmount,
			AssetName:         item.AssetName,
			Status:            item.Status,
			Tenor:             item.Tenor,
			CreatedAt:         item.CreatedAt,
			ApprovedAt:        item.ApprovedAt,
		})
	}

	return &payloads.GetLoanListResponse{Data: result, PaginationHelper: pagination}, err
}
