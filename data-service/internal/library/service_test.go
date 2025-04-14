package library_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
	"viska/data-service/internal/library"
	"viska/data-service/internal/library/mocks"
)

type LibrarySuite struct {
	suite.Suite
	ctx      context.Context
	svc      *library.LibraryService
	mockRepo *mocks.MockILibraryRepo
}

func (s *LibrarySuite) SetupSuite() {
	mockCtrl := gomock.NewController(s.T())
	defer mockCtrl.Finish()
	s.ctx = context.Background()
	s.mockRepo = mocks.NewMockILibraryRepo(mockCtrl)

	s.svc = library.NewLibraryService(s.mockRepo)
}

func TestRunServiceLibrarySuite(t *testing.T) {
	suite.Run(t, new(LibrarySuite))
}

func (s *LibrarySuite) TestCreateAuthorReturnSuccess() {
	var err error

	var author *library.Author
	author = &library.Author{Name: "J.K. Rowling"}

	s.mockRepo.EXPECT().CreateAuthor(gomock.Any()).Return(nil)

	err = s.svc.CreateAuthor(author)
	s.Assert().NoError(err)
}
