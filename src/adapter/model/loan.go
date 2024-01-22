package model

type LoanListRequest struct {
	UserId  int64  `json:"user_id"`
	Status  string `json:"status"`
	OrderBy string `json:"order_by"`
	CurPage int    `json:"cur_page"`
	Limit   int    `json:"limit"`
}
