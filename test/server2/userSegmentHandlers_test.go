package server2_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"testing"
	"time"

	"github.com/Bionic2113/avito/cmd/server"
	"github.com/Bionic2113/avito/internal/models"
	"github.com/stretchr/testify/require"
)

func TestUpdateSegmnetsForUser(t *testing.T) {
	var user_id uint64 = 1
	activeArr := []models.UserSegment{
		{Id: 1, User_id: user_id, Segment_name: "AVITO_TEST", CreationTime: uint64(time.Now().Unix()), Active: true},
		{Id: 2, User_id: user_id, Segment_name: "AVITO_MESSAGE", CreationTime: uint64(time.Now().Unix()), Active: true},
		{Id: 3, User_id: user_id, Segment_name: "OTIVA_VOICE", CreationTime: uint64(time.Now().Unix()), Active: true},
		{Id: 4, User_id: user_id, Segment_name: "OTIVA_TEST", CreationTime: uint64(time.Now().Unix()), Active: true},
		{Id: 5, User_id: user_id, Segment_name: "AVITO_VOICE", CreationTime: uint64(time.Now().Unix()), Active: true},
		{Id: 6, User_id: user_id, Segment_name: "AVITO_WRITE", CreationTime: uint64(time.Now().Unix()), Active: true},
	}
	deleteArr := []models.UserSegment{
		{
			Id:           7,
			User_id:      user_id,
			Segment_name: "AV_DEL",
			CreationTime: uint64(time.Now().Add(-1 * time.Hour).Unix()),
			DeletionTime: uint64(time.Date(2023, time.April, 30, 0, 0, 0, 0, time.Local).Unix()),
			Active:       false,
		},
		{
			Id:           8,
			User_id:      user_id,
			Segment_name: "AV_MEME",
			CreationTime: uint64(time.Now().Add(-1 * time.Hour).Unix()),
			DeletionTime: uint64(time.Date(2023, time.April, 30, 0, 0, 0, 0, time.Local).Unix()),
			Active:       false,
		},
		{
			Id:           9,
			User_id:      user_id,
			Segment_name: "OO_MEME",
			CreationTime: uint64(time.Now().Add(-1 * time.Hour).Unix()),
			DeletionTime: uint64(time.Date(2023, time.April, 30, 0, 0, 0, 0, time.Local).Unix()),
			Active:       false,
		},
		{
			Id:           10,
			User_id:      user_id,
			Segment_name: "AVITO_CLIP",
			CreationTime: uint64(time.Now().Add(-1 * time.Hour).Unix()),
			DeletionTime: uint64(time.Date(2023, time.April, 30, 0, 0, 0, 0, time.Local).Unix()),
			Active:       false,
		},
		{
			Id:           11,
			User_id:      user_id,
			Segment_name: "FAKE",
			CreationTime: uint64(time.Now().Add(-1 * time.Hour).Unix()),
			DeletionTime: uint64(time.Date(2023, time.April, 30, 0, 0, 0, 0, time.Local).Unix()),
			Active:       false,
		},
		{
			Id:           12,
			User_id:      user_id,
			Segment_name: "AVITO_LOTTERY",
			CreationTime: uint64(time.Now().Add(-1 * time.Hour).Unix()),
			DeletionTime: uint64(time.Date(2023, time.April, 30, 0, 0, 0, 0, time.Local).Unix()),
			Active:       false,
		},
		{
			Id:           13,
			User_id:      user_id,
			Segment_name: "AVITO_CASINO",
			CreationTime: uint64(time.Now().Add(-1 * time.Hour).Unix()),
			DeletionTime: uint64(time.Date(2023, time.April, 30, 0, 0, 0, 0, time.Local).Unix()),
			Active:       false,
		},
	}
	loadData(activeArr, deleteArr)

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
		models.UserSegment{Id: 14, User_id: user_id, Segment_name: "AVITO_NEW", Active: true},
		models.UserSegment{Id: 15, User_id: user_id, Segment_name: "AVITO_ADD", Active: true},
		models.UserSegment{Id: 16, User_id: user_id, Segment_name: "AVITO_HIP-HOP", Active: true},
	)

	request := server.UserSegmentsData{Id: int(user_id), Add: add, Del: del, Durations: []uint64{1, 2, 0}}

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
	excepted[13].CreationTime = actual[13].CreationTime
	excepted[14].CreationTime = actual[14].CreationTime
	excepted[15].CreationTime = actual[15].CreationTime
	excepted[13].Duration = actual[13].Duration
	excepted[14].Duration = actual[14].Duration
	excepted[15].Duration = actual[15].Duration

	require.Equal(t, excepted, actual, "Not equal. First test in TestUpdateSegmnetsForUser")
}

func loadData(activeArr []models.UserSegment, deletedArr []models.UserSegment) error {
	for _, value := range activeArr {
		_, err := segmentRepo.Create(value.Segment_name)
		if err != nil {
			log.Println("Create segments: ", err)
			return err
		}
	}
	u, err := userRepo.Create(&models.User{LastName: "Brown", FirstName: "David"})
	if err != nil {
		return err
	}
	log.Println("user: ", u.Id, u.FirstName, u.LastName, u.Active)
	for _, val := range activeArr {
		_, err := userSegmentRepo.Create(&val)
		if err != nil {
			log.Println("Create usersegments active: ", err)
			return err
		}
	}
	for _, val := range activeArr {
		_, err := userSegmentRepo.Update(&val)
		if err != nil {
			log.Println("update active: ", err)
			return err
		}
	}
	for _, value := range deletedArr {
		_, err := segmentRepo.Create(value.Segment_name)
		if err != nil {
			log.Println("create segments for delete: ", err)
			return nil
		}
	}
	for _, val := range deletedArr {
		_, err := userSegmentRepo.Create(&val)
		if err != nil {
			log.Println("create usersegments delete", err)
			return err
		}
	}
	for _, val := range deletedArr {
		_, err := userSegmentRepo.Update(&val)
		if err != nil {
			log.Println("update delete: ", err)
			return err
		}
	}

	return nil
}

// тест на автивных будет использовать то,
// что в предыдущем тесте уже появились данные
// это плохо, но время поджимает. Прошу прощения
func TestActiveUserSegmentsHandler(t *testing.T) {
	var user_id uint64 = 1
	excepted := []models.UserSegment{
		{Id: 2, User_id: user_id, Segment_name: "AVITO_MESSAGE", CreationTime: uint64(time.Now().Unix()), Active: true},
		{Id: 3, User_id: user_id, Segment_name: "OTIVA_VOICE", CreationTime: uint64(time.Now().Unix()), Active: true},
		{Id: 5, User_id: user_id, Segment_name: "AVITO_VOICE", CreationTime: uint64(time.Now().Unix()), Active: true},
		{Id: 14, User_id: user_id, Segment_name: "AVITO_NEW", CreationTime: uint64(time.Now().Unix()), Active: true},
		{Id: 15, User_id: user_id, Segment_name: "AVITO_ADD", CreationTime: uint64(time.Now().Unix()), Active: true},
		{Id: 16, User_id: user_id, Segment_name: "AVITO_HIP-HOP", CreationTime: uint64(time.Now().Unix()), Active: true},
	}
	request := &server.Message{Id: int(user_id)}
	req, _ := json.Marshal(request)
	resp, err := http.Post("http://localhost:8081/user_segment/active", "application/json", bytes.NewBuffer(req))
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
	sort.Slice(actual, func(i, j int) bool { return actual[i].Id < actual[j].Id })
	for i, v := range actual {
		excepted[i].CreationTime = v.CreationTime
		excepted[i].Duration = v.Duration
	}
	require.Equal(t, excepted, actual, "Not equal. First test in TestActiveUserSegmentsHandler")
}

func TestHistorySegmentsForUserHandler(t *testing.T) {
	request := &server.Date{Year: 2023, Month: 8}
	req, _ := json.Marshal(request)
	resp, err := http.Post("http://localhost:8081/user_segment/history", "application/json", bytes.NewBuffer(req))
	if err != nil {
		t.Log("Error in POST method")
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var actual string
	err = json.NewDecoder(resp.Body).Decode(&actual)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Link is : ", actual)
	require.NotEqual(t, "Not found elements", actual)

	request.Year = 2022
	req, _ = json.Marshal(request)
	resp, err = http.Post("http://localhost:8081/user_segment/history", "application/json", bytes.NewBuffer(req))
	if err != nil {
		t.Log("Error in POST method")
		t.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&actual)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Link is : ", actual)
	require.Equal(t, "Not found elements", actual)
}
