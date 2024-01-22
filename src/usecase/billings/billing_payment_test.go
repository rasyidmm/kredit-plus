package billings

import (
	"context"
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/mock"
	"kredit-plus/src/shared/utils/constant"
	"kredit-plus/src/usecase/payloads"
	"testing"
	"time"
)

func TestBillingsUsecase_BillingPayment(t *testing.T) {
	mockRepo := &mock.MockBillingRepository{}
	mockRepoLoan := &mock.MockLoansRepository{}
	Convey("Test Usecase BillingPayment", t, func() {
		Convey("Positive Scenarion", func() {
			Convey("BillingPayment Success", func() {
				mockRepoLoan.On("GetLoanById", context.Background(), int64(0)).Return(&entity.LoansEntity{}, nil).Once()
				mockRepo.On("GetBillingListByLoanIdAndStatus", context.Background(), int64(0), []string{constant.InReview, constant.Approved}).Return(&[]entity.BillingsEntity{}, nil).Once()
				mockRepo.On("CreateBilling", context.Background(), &entity.BillingsEntity{CreatedAt: time.Now(), LoanId: int64(0), Status: constant.InReview}).Return(nil).Once()
				uc := NewBillingsUsecase(mockRepo, mockRepoLoan)
				err := uc.BillingPayment(context.Background(), &payloads.BillingPaymentRequest{LoanId: 0})
				So(err, ShouldBeNil)
			})
		})
		Convey("Negative Scenarion", func() {
			Convey("BillingPayment Failed", func() {
				mockRepoLoan.On("GetLoanById", context.Background(), int64(0)).Return(nil, errors.New("error")).Once()
				mockRepo.On("GetBillingListByLoanIdAndStatus", context.Background(), int64(0), []string{constant.InReview, constant.Approved}).Return(nil, errors.New("error")).Once()
				mockRepo.On("CreateBilling", context.Background(), &entity.BillingsEntity{CreatedAt: time.Now(), LoanId: int64(0), Status: constant.InReview}).Return(errors.New("error")).Once()
				uc := NewBillingsUsecase(mockRepo, mockRepoLoan)
				err := uc.BillingPayment(context.Background(), &payloads.BillingPaymentRequest{LoanId: 0})
				So(err, ShouldNotBeNil)
			})

		})
	})
}
