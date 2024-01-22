package payloads

import (
	"kredit-plus/src/shared/utils"
	"time"
)

type SubmissionLoanRequest struct {
	UserId    int64   `json:"user_id"`
	AdminFee  float64 `json:"admin_fee"`
	Tenor     int     `json:"tenor"`
	AssetName string  `json:"asset_name"`
	Amount    float64 `json:"amount"`
}
type LoanListRequest struct {
	UserId  int64  `json:"user_id"`
	Status  string `json:"status"`
	OrderBy string `json:"order_by"`
	CurPage int    `json:"cur_page"`
	Limit   int    `json:"limit"`
}

type LoanResponse struct {
	Id                int64      `json:"id"`
	UserId            int64      `json:"user_id"`
	LoanNo            string     `json:"loan_no"`
	Otr               float64    `json:"otr"`
	AdminFee          float64    `json:"admin_fee"`
	InstallmentAmount float64    `json:"installment_amount"`
	InterestAmount    float64    `json:"interest_amount"`
	AssetName         string     `json:"asset_name"`
	Status            string     `json:"status"`
	Tenor             int        `json:"tenor"`
	CreatedAt         time.Time  `json:"created_at"`
	ApprovedAt        *time.Time `json:"approved_at"`
}

type GetLoanListResponse struct {
	Data             []LoanResponse          `json:"data"`
	PaginationHelper *utils.PaginationHelper `json:"pagination"`
}
