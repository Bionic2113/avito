package server_test

import (
	"testing"

	"github.com/Bionic2113/avito/cmd/server"
	"github.com/Bionic2113/avito/internal/repository"
	database_test "github.com/Bionic2113/avito/test/database"
	"github.com/Bionic2113/avito/pkg/fakedb"

	gomock "go.uber.org/mock/gomock"
)

type Message struct {
	Id int `json:"id"`
}

var (
	userMockDB    *database_test.MockUserRepository
	segmentMockDB *database_test.MockSegmentRepository
	usMockDB      *database_test.MockUserSegmentRepository

	userRepo        repository.UserRepository
	segmentRepo     repository.SegmentRepository
	userSegmentRepo repository.UserSegmentRepository
)

func TestMain(m *testing.M) {
	t := &testing.T{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userMockDB = database_test.NewMockUserRepository(ctrl)
	segmentMockDB = database_test.NewMockSegmentRepository(ctrl)
	usMockDB = database_test.NewMockUserSegmentRepository(ctrl)
	go server.StartServer(userMockDB, segmentMockDB, usMockDB)

	db, err := fakedb.SetupDB()
	if err != nil {
		t.Fatal("Не создалась бд", err)
	}

	// подруби тут взятие данных из файла
	// db, err := sql.Open("postgres", "user=postgres password=123 dbname=test_db sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	userRepo = &repository.UserRepoDB{DB: db}
	segmentRepo = &repository.SegmentRepositoryDB{DB: db}
	userSegmentRepo = &repository.UserSegmentDB{DB: db}
	server.Port = 8081
	go server.StartServer(userRepo, segmentRepo, userSegmentRepo)

	m.Run()

	if err = fakedb.DropDB(db); err != nil {
		t.Fatal("Не удалилась бд", err)
	}
}
