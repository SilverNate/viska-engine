package authentication_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
	"viska/auth-service/config"
	auth "viska/auth-service/internal/authentication"
	"viska/auth-service/internal/authentication/mocks"
)

type AuthSuite struct {
	suite.Suite
	ctx      context.Context
	svc      *auth.AuthServiceImpl
	mockRepo *mocks.MockAuthRepository
	config   config.Config
}

func (s *AuthSuite) SetupSuite() {
	mockCtrl := gomock.NewController(s.T())
	defer mockCtrl.Finish()
	s.ctx = context.Background()
	s.mockRepo = mocks.NewMockAuthRepository(mockCtrl)

	s.svc = auth.NewAuthService(config.Config{}, s.mockRepo)
}

func TestRunServiceAuthSuite(t *testing.T) {
	suite.Run(t, new(AuthSuite))
}

func (s *AuthSuite) TestGetUserByEmailReturnSuccess() {
	email := "admin@viska.io"
	var err error

	var users *auth.User
	users = &auth.User{
		Email: email,
	}

	expectedResult := users

	s.mockRepo.EXPECT().GetUserByEmail(gomock.Any()).Return(users, nil)

	actualResult, err := s.svc.GetUserByEmail(email)
	s.Assert().NoError(err)
	s.Assert().Equal(expectedResult, actualResult)
}
