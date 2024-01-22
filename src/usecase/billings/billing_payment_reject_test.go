package billings

import (
	"context"
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/mock"
	"kredit-plus/src/shared/utils/constant"
	"testing"
)

func TestBillingsUsecase_BillingPaymentReject(t *testing.T) {
	mockRepo := &mock.MockBillingRepository{}
	mockRepoLoan := &mock.MockLoansRepository{}
	Convey("Test Usecase BillingPaymentReject", t, func() {
		Convey("Positive Scenarion", func() {
			Convey("BillingPaymentReject Success", func() {
				mockRepo.On("GetBillingById", context.Background(), int64(1)).Return(&entity.BillingsEntity{}, nil).Once()
				mockRepo.On("UpdateBilling", context.Background(), &entity.BillingsEntity{Status: constant.Rejected}).Return(nil).Once()
				uc := NewBillingsUsecase(mockRepo, mockRepoLoan)
				err := uc.BillingPaymentReject(context.Background(), int64(1))
				So(err, ShouldBeNil)
			})
		})
		Convey("Negative Scenarion", func() {
			Convey("BillingPaymentReject Failed", func() {
				mockRepo.On("GetBillingById", context.Background(), int64(1)).Return(&entity.BillingsEntity{}, errors.New("error")).Once()
				mockRepo.On("UpdateBilling", context.Background(), &entity.BillingsEntity{Status: constant.Rejected}).Return(errors.New("error")).Once()
				uc := NewBillingsUsecase(mockRepo, mockRepoLoan)
				err := uc.BillingPaymentReject(context.Background(), int64(1))
				So(err, ShouldNotBeNil)
			})

		})
	})
}
