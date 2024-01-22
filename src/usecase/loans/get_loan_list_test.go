package loans

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

func TestLoansUsecase_GetLoanList(t *testing.T) {
	mockRepo := &mock.MockLoansRepository{}
	mockRepoTenor := &mock.MockMasterTenorRepository{}
	MockRepoUser := &mock.MockUsersRepository{}
	Convey("Test Usecase GetLoanList", t, func() {
		Convey("Positive Scenarion", func() {
			Convey("GetLoanById Success", func() {
				mockRepo.On("GetLoanList", context.Background(), &model.LoanListRequest{}).Return(&[]entity.LoansEntity{}, &utils.PaginationHelper{}, nil).Once()
				uc := NewLoansUsecase(mockRepo, mockRepoTenor, MockRepoUser)
				res, err := uc.GetLoanList(context.Background(), &payloads.LoanListRequest{})
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
		Convey("Negative Scenarion", func() {
			Convey("GetLoanList Failed", func() {
				mockRepo.On("GetLoanList", context.Background(), &model.LoanListRequest{}).Return(nil, &utils.PaginationHelper{}, errors.New("error")).Once()
				uc := NewLoansUsecase(mockRepo, mockRepoTenor, MockRepoUser)
				res, err := uc.GetLoanList(context.Background(), &payloads.LoanListRequest{})
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})

		})
	})
}
