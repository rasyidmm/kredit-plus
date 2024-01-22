package payloads

type MasterTenor struct {
	Id       *int64  `json:"id"`
	UserId   int64   `json:"user_id"`
	Tenor    int64   `json:"tenor"`
	Amount   float64 `json:"amount"`
	Interest float64 `json:"interest"`
}
