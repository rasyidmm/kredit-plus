package loans

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

func TestLoansUsecase_SubmissionLoans(t *testing.T) {
	mockRepo := &mock.MockLoansRepository{}
	mockRepoTenor := &mock.MockMasterTenorRepository{}
	MockRepoUser := &mock.MockUsersRepository{}
	Convey("Test Usecase SubmissionLoans", t, func() {
		Convey("Positive Scenarion", func() {
			Convey("SubmissionLoans Success", func() {
				MockRepoUser.On("GetUserById", context.Background(), int64(0)).Return(&entity.UsersEntity{}, nil).Once()
				mockRepoTenor.On("GetMasterTenorByUserIdAndTenor", context.Background(), int64(0), 0).Return(&entity.MasterTenorEntity{}, nil).Once()
				mockRepo.On("GetLoanListByUserIdAndTenor", context.Background(), int64(0), 0, []string{constant.InReview, constant.Approved}).Return(&[]entity.LoansEntity{}, nil).Once()
				format := "_" + time.Now().Format("20060102") + "_" + time.Now().Format("150405")
				mockRepo.On("CreateLoan", context.Background(), &entity.LoansEntity{CreatedAt: time.Now(), Status: constant.InReview, LoanNo: format}).Return(nil).Once()
				uc := NewLoansUsecase(mockRepo, mockRepoTenor, MockRepoUser)
				err := uc.SubmissionLoans(context.Background(), &payloads.SubmissionLoanRequest{})
				So(err, ShouldBeNil)
			})
		})
		Convey("Negative Scenarion", func() {
			Convey("SubmissionLoans Failed", func() {
				MockRepoUser.On("GetUserById", context.Background(), int64(0)).Return(&entity.UsersEntity{}, nil).Once()
				mockRepoTenor.On("GetMasterTenorByUserIdAndTenor", context.Background(), int64(0), 0).Return(nil, errors.New("error")).Once()
				mockRepo.On("GetLoanListByUserIdAndTenor", context.Background(), int64(0), 0, []string{constant.InReview, constant.Approved}).Return(nil, errors.New("error")).Once()
				format := "_" + time.Now().Format("20060102") + "_" + time.Now().Format("150405")
				mockRepo.On("CreateLoan", context.Background(), &entity.LoansEntity{CreatedAt: time.Now(), Status: constant.InReview, LoanNo: format}).Return(errors.New("error")).Once()
				uc := NewLoansUsecase(mockRepo, mockRepoTenor, MockRepoUser)
				err := uc.SubmissionLoans(context.Background(), &payloads.SubmissionLoanRequest{})
				So(err, ShouldNotBeNil)

			})

		})
	})
}
