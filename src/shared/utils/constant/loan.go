package constant

var StatusLoanOrderBy = map[string]bool{
	Rejected: true,
	InReview: true,
	Approved: true,
}
var LoanIsOrderBy = map[string]bool{
	"id":          true,
	"created_at":  true,
	"tenor":       true,
	"approved_at": true,
}
