package loans

import (
	"context"
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"kredit-plus/src/adapter/repository/entity"
	"kredit-plus/src/shared/mock"
	"testing"
)

func TestLoansUsecase_GetLoanById(t *testing.T) {
	mockRepo := &mock.MockLoansRepository{}
	mockRepoTenor := &mock.MockMasterTenorRepository{}
	MockRepoUser := &mock.MockUsersRepository{}
	Convey("Test Usecase GetLoanById", t, func() {
		Convey("Positive Scenarion", func() {
			Convey("GetLoanById Success", func() {
				mockRepo.On("GetLoanById", context.Background(), int64(1)).Return(&entity.LoansEntity{}, nil).Once()
				uc := NewLoansUsecase(mockRepo, mockRepoTenor, MockRepoUser)
				res, err := uc.GetLoanById(context.Background(), int64(1))
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
		Convey("Negative Scenarion", func() {
			Convey("GetLoanById Failed", func() {
				mockRepo.On("GetLoanById", context.Background(), int64(1)).Return(nil, errors.New("error")).Once()
				uc := NewLoansUsecase(mockRepo, mockRepoTenor, MockRepoUser)
				res, err := uc.GetLoanById(context.Background(), int64(1))
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})

		})
	})
}
