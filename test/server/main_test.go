package server_test

import (
	"testing"

	"github.com/Bionic2113/avito/cmd/server"
	database_test "github.com/Bionic2113/avito/test/database"

	gomock "go.uber.org/mock/gomock"
)

type Message struct {
	Id int `json:"id"`
}

var (
  userMockDB *database_test.MockUserRepository
  segmentMockDB *database_test.MockSegmentRepository
  usMockDB *database_test.MockUserSegmentRepository 
)

func TestMain(m *testing.M) {
	ctrl := gomock.NewController(&testing.T{})
	defer ctrl.Finish()
	userMockDB = database_test.NewMockUserRepository(ctrl)
	segmentMockDB = database_test.NewMockSegmentRepository(ctrl)
	go server.StartServer(userMockDB, segmentMockDB, usMockDB)
	m.Run()
}
