package server_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/Bionic2113/avito/cmd/server"
	"github.com/Bionic2113/avito/internal/models"
	"github.com/stretchr/testify/require"
)

func TestCreateSegmentHandler(t *testing.T) {
	segment_start := server.SegmentName{Name: "AVITO_TEST"}
	segment_finish := models.Segment{1, segment_start.Name, true}
	segmentMockDB.EXPECT().Create(segment_start.Name).Return(&segment_finish, nil)
	req, _ := json.Marshal(segment_start)
	resp, err := http.Post("http://localhost:8080/segment/create", "application/json", bytes.NewBuffer(req))
	if err != nil {
		t.Log("error in POST method")
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var segment_actual models.Segment
	err = json.NewDecoder(resp.Body).Decode(&segment_actual)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, segment_finish, segment_actual, "Not equal. First test in createSegmentHandler")
}

// почему то стал отлетать, хотя в постмане всё ок
func TestDeleteSegmentHandler(t *testing.T) {
	testTable := []struct {
		s        server.SegmentName
		err      error
		response string
	}{
		{s: server.SegmentName{Name: "AVITO_TEST_2"}, err: nil, response: "Success"},
		{s: server.SegmentName{}, err: errors.New("Not found"), response: "Error! Not found!"},
	}
	for i, v := range testTable {
		segmentMockDB.EXPECT().Delete(v.s.Name).Return(v.err)
		req, err := json.Marshal(v.s)
		if err != nil {
			t.Log("eror in marshal: ", err, "iter: ", i)
			t.Fatal(err)
		}
		resp, err := http.Post("http://localhost:8080/segment/delete", "application/json", bytes.NewBuffer(req))
		if err != nil {
			t.Log("error in POST method. Iteration = ", i)
			t.Fatal(err)
		}
		defer resp.Body.Close()

		var resp_actual string
		err = json.NewDecoder(resp.Body).Decode(&resp_actual)
		if err != nil {
			t.Log("second ", resp.Body)
			t.Log("error in decode, iteration = ", i)
			t.Fatal(err)
		}

		require.Equal(t, v.response, resp_actual, fmt.Sprintf("Not equal. Erorr in %d test deleteSegmentHandler", i))

	}
}
