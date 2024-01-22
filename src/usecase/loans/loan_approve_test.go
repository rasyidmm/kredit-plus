package loans

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

func TestLoansUsecase_LoanApprove(t *testing.T) {
	mockRepo := &mock.MockLoansRepository{}
	mockRepoTenor := &mock.MockMasterTenorRepository{}
	MockRepoUser := &mock.MockUsersRepository{}
	Convey("Test Usecase LoanApprove", t, func() {
		Convey("Positive Scenarion", func() {
			Convey("LoanApprove Success", func() {
				mockRepo.On("GetLoanById", context.Background(), int64(1)).Return(&entity.LoansEntity{}, nil).Once()
				tt := time.Now()
				mockRepo.On("UpdateLoan", context.Background(), &entity.LoansEntity{Status: constant.Approved, ApprovedAt: &tt}).Return(nil).Once()
				uc := NewLoansUsecase(mockRepo, mockRepoTenor, MockRepoUser)
				err := uc.LoanApprove(context.Background(), int64(1))
				So(err, ShouldBeNil)
			})
		})
		Convey("Negative Scenarion", func() {
			Convey("LoanApprove Failed", func() {
				mockRepo.On("GetLoanById", context.Background(), int64(1)).Return(nil, errors.New("error")).Once()
				tt := time.Now()
				mockRepo.On("UpdateLoan", context.Background(), &entity.LoansEntity{Status: constant.Approved, ApprovedAt: &tt}).Return(errors.New("error")).Once()
				uc := NewLoansUsecase(mockRepo, mockRepoTenor, MockRepoUser)
				err := uc.LoanApprove(context.Background(), int64(1))
				So(err, ShouldNotBeNil)

			})

		})
	})
}
