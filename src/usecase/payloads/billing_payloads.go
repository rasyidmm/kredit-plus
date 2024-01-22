package payloads

import (
	"kredit-plus/src/shared/utils"
	"time"
)

type BillingPaymentRequest struct {
	LoanId int64   `json:"loan_id"`
	Amount float64 `json:"amount"`
	UserId int64   `json:"user_id"`
}

type BillingListRequest struct {
	UserId  int64  `json:"user_id"`
	LoanId  int64  `json:"loan_id"`
	Status  string `json:"status"`
	OrderBy string `json:"order_by"`
	CurPage int    `json:"cur_page"`
	Limit   int    `json:"limit"`
}
type BillingResponse struct {
	Id         int64      `json:"id"`
	LoanId     int64      `json:"loan_id"`
	UserId     int64      `json:"user_id"`
	CreatedAt  time.Time  `json:"created_at"`
	Amount     float64    `json:"amount"`
	Status     string     `json:"status"`
	ApprovedAt *time.Time `json:"approved_at"`
}
type BillingListResponse struct {
	Data             []BillingResponse       `json:"data"`
	PaginationHelper *utils.PaginationHelper `json:"pagination"`
}
