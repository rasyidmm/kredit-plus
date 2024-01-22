package billings

import (
	"context"
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"kredit-plus/src/adapter/model"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/mock"
	"kredit-plus/src/shared/utils"
	"kredit-plus/src/usecase/payloads"
	"testing"
)

func TestBillingsUsecase_GetBillingList(t *testing.T) {
	mockRepo := &mock.MockBillingRepository{}
	mockRepoLoan := &mock.MockLoansRepository{}
	Convey("Test Usecase GetBillingList", t, func() {
		Convey("Positive Scenarion", func() {
			Convey("GetBillingList Success", func() {
				mockRepo.On("GetBillingList", context.Background(), &model.BillingListRequest{}).Return(&[]entity.BillingsEntity{}, &utils.PaginationHelper{}, nil).Once()
				uc := NewBillingsUsecase(mockRepo, mockRepoLoan)
				res, err := uc.GetBillingList(context.Background(), &payloads.BillingListRequest{})
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
		Convey("Negative Scenarion", func() {
			Convey("GetBillingList Failed", func() {
				mockRepo.On("GetBillingList", context.Background(), &model.BillingListRequest{}).Return(nil, &utils.PaginationHelper{}, errors.New("error")).Once()
				uc := NewBillingsUsecase(mockRepo, mockRepoLoan)
				res, err := uc.GetBillingList(context.Background(), &payloads.BillingListRequest{})
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})

		})
	})

}
