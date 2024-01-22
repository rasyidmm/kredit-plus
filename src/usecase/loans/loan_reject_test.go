package loans

import (
	"context"
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/mock"
	"kredit-plus/src/shared/utils/constant"
	"testing"
)

func TestLoansUsecase_LoanReject(t *testing.T) {
	mockRepo := &mock.MockLoansRepository{}
	mockRepoTenor := &mock.MockMasterTenorRepository{}
	MockRepoUser := &mock.MockUsersRepository{}
	Convey("Test Usecase LoanReject", t, func() {
		Convey("Positive Scenarion", func() {
			Convey("LoanReject Success", func() {
				mockRepo.On("GetLoanById", context.Background(), int64(1)).Return(&entity.LoansEntity{}, nil).Once()
				mockRepo.On("UpdateLoan", context.Background(), &entity.LoansEntity{Status: constant.Rejected}).Return(nil).Once()
				uc := NewLoansUsecase(mockRepo, mockRepoTenor, MockRepoUser)
				err := uc.LoanReject(context.Background(), int64(1))
				So(err, ShouldBeNil)
			})
		})
		Convey("Negative Scenarion", func() {
			Convey("LoanReject Failed", func() {
				mockRepo.On("GetLoanById", context.Background(), int64(1)).Return(nil, errors.New("error")).Once()
				mockRepo.On("UpdateLoan", context.Background(), &entity.LoansEntity{Status: constant.Rejected}).Return(errors.New("error")).Once()
				uc := NewLoansUsecase(mockRepo, mockRepoTenor, MockRepoUser)
				err := uc.LoanReject(context.Background(), int64(1))
				So(err, ShouldNotBeNil)

			})

		})
	})
}
