package constant

var (
	Approved = "APPROVED"
	InReview = "IN_REVIEW"
	Rejected = "REJECTED"
	Done     = "DONE"
)

var UserIsOrderBy = map[string]bool{
	"id":         true,
	"created_at": true,
}
var OrderAction = map[string]bool{
	"asc":  true,
	"desc": true,
}

var StatusUserOrderBy = map[string]bool{
	Rejected: true,
	InReview: true,
	Approved: true,
	Done:     true,
}
