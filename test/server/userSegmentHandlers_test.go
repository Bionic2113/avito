package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/Bionic2113/avito/cmd/server"
	"github.com/Bionic2113/avito/internal/models"
	"github.com/stretchr/testify/require"
)

func TestUpdateSegmnetsForUser(t *testing.T) {
	activeArr := []models.UserSegment{
		{Id: 1, User_id: 3, Segment_name: "AVITO_TEST", CreationTime: uint64(time.Now().Unix()), Active: true},
		{Id: 2, User_id: 3, Segment_name: "AVITO_MESSAGE", CreationTime: uint64(time.Now().Unix()), Active: true},
		{Id: 3, User_id: 3, Segment_name: "OTIVA_VOICE", CreationTime: uint64(time.Now().Unix()), Active: true},
		{Id: 4, User_id: 3, Segment_name: "OTIVA_TEST", CreationTime: uint64(time.Now().Unix()), Active: true},
		{Id: 5, User_id: 3, Segment_name: "AVITO_VOICE", CreationTime: uint64(time.Now().Unix()), Active: true},
		{Id: 6, User_id: 3, Segment_name: "AVITO_WRITE", CreationTime: uint64(time.Now().Unix()), Active: true},
	}
	deleteArr := []models.UserSegment{
		{
			Id:           7,
			User_id:      3,
			Segment_name: "AV_DEL",
			DeletionTime: uint64(time.Date(2023, time.April, 30, 0, 0, 0, 0, time.Local).Unix()),
			Active:       false,
		},
		{
			Id:           8,
			User_id:      3,
			Segment_name: "AV_MEME",
			CreationTime: uint64(time.Now().Add(-1 * time.Hour).Unix()),
			DeletionTime: uint64(time.Date(2023, time.April, 30, 0, 0, 0, 0, time.Local).Unix()),
			Active:       false,
		},
		{
			Id:           9,
			User_id:      3,
			Segment_name: "OO_MEME",
			CreationTime: uint64(time.Now().Add(-1 * time.Hour).Unix()),
			DeletionTime: uint64(time.Date(2023, time.April, 30, 0, 0, 0, 0, time.Local).Unix()),
			Active:       false,
		},
		{
			Id:           10,
			User_id:      3,
			Segment_name: "AVITO_CLIP",
			CreationTime: uint64(time.Now().Add(-1 * time.Hour).Unix()),
			DeletionTime: uint64(time.Date(2023, time.April, 30, 0, 0, 0, 0, time.Local).Unix()),
			Active:       false,
		},
		{
			Id:           11,
			User_id:      3,
			Segment_name: "FAKE",
			CreationTime: uint64(time.Now().Add(-1 * time.Hour).Unix()),
			DeletionTime: uint64(time.Date(2023, time.April, 30, 0, 0, 0, 0, time.Local).Unix()),
			Active:       false,
		},
		{
			Id:           12,
			User_id:      3,
			Segment_name: "AVITO_LOTTERY",
			CreationTime: uint64(time.Now().Add(-1 * time.Hour).Unix()),
			DeletionTime: uint64(time.Date(2023, time.April, 30, 0, 0, 0, 0, time.Local).Unix()),
			Active:       false,
		},
		{
			Id:           14,
			User_id:      3,
			Segment_name: "AVITO_CASINO",
			CreationTime: uint64(time.Now().Add(-1 * time.Hour).Unix()),
			DeletionTime: uint64(time.Date(2023, time.April, 30, 0, 0, 0, 0, time.Local).Unix()),
			Active:       false,
		},
	}

	// usMockDB.EXPECT().FindAllById(3, true).Return(activeArr, nil)
	// usMockDB.EXPECT().FindAllById(3, false).Return(deleteArr, nil)

	// userMockDB.EXPECT().FindById(3).Return(&models.User{Id:3}, nil)

	add := []string{
		"AVITO_NEW",
		"AVITO_ADD",
		"AVITO_HIP-HOP",
	}
	del := []string{
		"AVITO_TEST",
		"OTIVA_TEST",
		"AVITO_WRITE",
	}
	// сначала удаляем, потом добавляем старые удаленные, а потом уже новые
	excepted := activeArr
	excepted[0].Active = false
	excepted[3].Active = false
	excepted[5].Active = false
	excepted = append(excepted, deleteArr...)
	excepted = append(
	  excepted, 
	  models.UserSegment{Id: 15, User_id: 3, Segment_name: "AVITO_NEW"},
	  models.UserSegment{Id: 16, User_id: 3, Segment_name: "AVITO_ADD"},
	  models.UserSegment{Id: 17, User_id: 3, Segment_name: "AVITO_HIP-HOP"},
	)

	request := server.UserSegmentsData{Id: 3, Add: add, Del: del}

	req, _ := json.Marshal(request)
	resp, err := http.Post("http://localhost:8081/user_segment/update", "application/json", bytes.NewBuffer(req))
	if err != nil {
		t.Log("Error in POST method")
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var actual []models.UserSegment
	err = json.NewDecoder(resp.Body).Decode(&actual)
	if err != nil {
		t.Fatal(err)
	}
	// так время удаление берется нынешнее, то его невозможно задать самостоятельно
	// поэтому я его просто возьму из полученного результа, тк еще
	// есть active, который можно проверить на false 
	excepted[0].DeletionTime = actual[0].DeletionTime
	excepted[3].DeletionTime = actual[3].DeletionTime
	excepted[5].DeletionTime = actual[5].DeletionTime
	// в создании время тоже само создается базой данных,
	// тк я поставил default = now(), поэтому снова берем его
	excepted[14].CreationTime = actual[14].CreationTime
	excepted[15].CreationTime = actual[15].CreationTime
	excepted[16].CreationTime = actual[16].CreationTime


	require.Equal(t, excepted, actual, "Not equal. First test in TestUpdateSegmnetsForUser")
}

func loadData(user_id int, activeArr []models.UserSegment, deletedArr []models.UserSegment) error{
  for _, value := range activeArr{
	_, err := segmentRepo.Create(value.Segment_name)
	if err != nil{
	  return err
	}
  }
  _, err := userRepo.Create(&models.User{LastName:"Brown", FirstName:"David"})
  if err != nil{
	return err
  }
  for _, val := range activeArr{
	_, err := userSegmentRepo.Create(&val)
	if err != nil{return err}
  }
  for _, value := range deletedArr{
	_, err := segmentRepo.Create(value.Segment_name)
	if err != nil{return nil}
  }
  for _, val := range deletedArr{
	_, err := userSegmentRepo.Create(&val)
	if err != nil {return err}
  }
  return nil
}
