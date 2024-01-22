package model

type GetUserListRequest struct {
	Status  string `json:"status"`
	OrderBy string `json:"order_by"`
	CurPage int    `json:"cur_page"`
	Limit   int    `json:"limit"`
}
