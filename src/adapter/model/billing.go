package model

type BillingListRequest struct {
	UserId  int64  `json:"user_id"`
	LoanId  int64  `json:"loan_id"`
	Status  string `json:"status"`
	OrderBy string `json:"order_by"`
	CurPage int    `json:"cur_page"`
	Limit   int    `json:"limit"`
}
