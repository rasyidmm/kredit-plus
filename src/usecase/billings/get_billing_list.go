package billings

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"kredit-plus/src/adapter/model"
	"kredit-plus/src/usecase/payloads"
	"net/http"
)

func (u *BillingsUsecase) GetBillingList(ctx context.Context, req *payloads.BillingListRequest) (*payloads.BillingListResponse, error) {

	data, pagination, err := u.repo.GetBillingList(ctx, &model.BillingListRequest{
		UserId:  req.UserId,
		LoanId:  req.LoanId,
		Status:  req.Status,
		OrderBy: req.OrderBy,
		CurPage: req.CurPage,
		Limit:   req.Limit,
	})
	if err != nil {
		return nil, fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	var result []payloads.BillingResponse

	for _, item := range *data {
		result = append(result, payloads.BillingResponse{
			Id:         item.Id,
			LoanId:     item.LoanId,
			UserId:     item.UserId,
			CreatedAt:  item.CreatedAt,
			Amount:     item.Amount,
			Status:     item.Status,
			ApprovedAt: item.ApprovedAt,
		})
	}
	return &payloads.BillingListResponse{
		Data:             result,
		PaginationHelper: pagination,
	}, err
}
