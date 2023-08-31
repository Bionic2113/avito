package server3_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type SegP struct {
	SegmentName string `json:"segment_name"`
	Percent     int    `json:"percent"`
}

// Проверка осуществляется по количеству
// добавленных нужных записей, тк
// сравнивать нет смысла
func TestAddPercent(t *testing.T) {
	testTable := []struct {
		sp     SegP
		status bool
	}{
		{sp: SegP{"TESTIFY", 40}, status: true},
		{sp: SegP{"MUMUMU", 4000}, status: false},
	}
	for _, v := range testTable {
		u_arr, err := userSegmentRepo.FindAll()
		if err != nil {
			t.Fatal(err)
		}
		var count_1 int
		for _, value := range u_arr {
			if v.sp.SegmentName == value.Segment_name {
				count_1++
			}
		}
		req, _ := json.Marshal(v.sp)
		resp, err := http.Post("http://localhost:8082/user_segment/add_percent", "application/json", bytes.NewBuffer(req))
		if err != nil {
			t.Log("error in POST method")
			t.Fatal(err)
		}
		defer resp.Body.Close()

		// var segment_actual models.Segment
		// err = json.NewDecoder(resp.Body).Decode(&segment_actual)
		// if err != nil {
		// 	t.Fatal(err)
		// }
		u_arr_2, err := userSegmentRepo.FindAll()
		if err != nil {
			t.Fatal(err)
		}
		var count_2 int
		for _, value := range u_arr_2 {
			t.Log(v.sp.SegmentName," ",value.Segment_name)
			if v.sp.SegmentName == value.Segment_name {
				count_2++
			}
		}
		t.Log("Segment_name: ", v.sp.SegmentName, " status: ", v.status, " count_1: ", count_1, " count_2: ", count_2)
		require.Equal(t, v.status, count_1 != count_2)
	}
}
