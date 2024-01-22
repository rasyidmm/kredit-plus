package billings

import (
	"context"
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/mock"
	"kredit-plus/src/shared/utils/constant"
	"testing"
	"time"
)

func TestBillingsUsecase_BillingPaymentApprove(t *testing.T) {
	mockRepo := &mock.MockBillingRepository{}
	mockRepoLoan := &mock.MockLoansRepository{}
	Convey("Test Usecase BillingPaymentApprove", t, func() {
		Convey("Positive Scenarion", func() {
			Convey("BillingPaymentApprove Success", func() {

				mockRepo.On("GetBillingById", context.Background(), int64(1)).Return(&entity.BillingsEntity{}, nil).Once()
				tt := time.Now()
				mockRepo.On("UpdateBilling", context.Background(), &entity.BillingsEntity{Status: constant.Approved, ApprovedAt: &tt}).Return(nil).Once()
				uc := NewBillingsUsecase(mockRepo, mockRepoLoan)
				err := uc.BillingPaymentApprove(context.Background(), int64(1))
				So(err, ShouldBeNil)
			})
		})
		Convey("Negative Scenarion", func() {
			Convey("BillingPaymentApprove Failed", func() {

				mockRepo.On("GetBillingById", context.Background(), int64(1)).Return(&entity.BillingsEntity{}, errors.New("error")).Once()
				tt := time.Now()
				mockRepo.On("UpdateBilling", context.Background(), &entity.BillingsEntity{Status: constant.Approved, ApprovedAt: &tt}).Return(errors.New("error")).Once()
				uc := NewBillingsUsecase(mockRepo, mockRepoLoan)
				err := uc.BillingPaymentApprove(context.Background(), int64(1))
				So(err, ShouldNotBeNil)
			})

		})
	})
}
