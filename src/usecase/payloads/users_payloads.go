package payloads

import (
	"kredit-plus/src/shared/utils"
	"time"
)

type RegistrationRequest struct {
	Nik                 string  `json:"nik"`
	FullName            string  `json:"full_name"`
	LegalName           string  `json:"legal_name"`
	Birthplace          string  `json:"birthplace"`
	Birthdate           string  `json:"birthdate"`
	Salary              float64 `json:"salary"`
	IdentificationPhoto string  `json:"identification_photo"`
	PhotoSelfie         string  `json:"photo_selfie"`
}

type GetUserListRequest struct {
	Status  string `json:"status"`
	OrderBy string `json:"order_by"`
	CurPage int    `json:"cur_page"`
	Limit   int    `json:"limit"`
}

type UserResponse struct {
	Id                  *int64
	Nik                 string
	FullName            string
	LegalName           string
	Birthplace          string
	Birthdate           time.Time
	Salary              float64
	IdentificationPhoto string
	PhotoSelfie         string
	Status              string
	CreatedAt           time.Time
	Tenor               *[]MasterTenor
}

type GetUserListResponse struct {
	Data             []UserResponse          `json:"data"`
	PaginationHelper *utils.PaginationHelper `json:"pagination"`
}
