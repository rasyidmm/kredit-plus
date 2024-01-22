package billings

import (
	"context"
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/mock"
	"testing"
)

func TestBillingsUsecase_GetBillingById(t *testing.T) {
	mockRepo := &mock.MockBillingRepository{}
	mockRepoLoan := &mock.MockLoansRepository{}
	Convey("Test Usecase GetBillingById", t, func() {
		Convey("Positive Scenarion", func() {
			Convey("Get Billing By Id Success", func() {
				mockRepo.On("GetBillingById", context.Background(), int64(1)).Return(&entity.BillingsEntity{}, nil).Once()
				uc := NewBillingsUsecase(mockRepo, mockRepoLoan)
				res, err := uc.GetBillingById(context.Background(), int64(1))
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
		Convey("Negative Scenarion", func() {
			Convey("Get Billing By Id Failed", func() {
				mockRepo.On("GetBillingById", context.Background(), int64(1)).Return(nil, errors.New("error")).Once()
				uc := NewBillingsUsecase(mockRepo, mockRepoLoan)
				res, err := uc.GetBillingById(context.Background(), 1)
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})

		})
	})
}
