package server_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/Bionic2113/avito/internal/models"
	"github.com/stretchr/testify/assert"
)

type SegmentName struct {
	Name string `json:"name,omitempty"`
}

func TestCreateSegmentHandler(t *testing.T) {
	segment_start := SegmentName{Name: "AVITO_TEST"}
	segment_finish := models.Segment{1, segment_start.Name, true}
	segmentMockDB.EXPECT().Create(segment_start.Name).Return(&segment_finish, nil)
	req, _ := json.Marshal(segment_start)
	resp, err := http.Post("http://localhost:8080/segment/create", "application/json", bytes.NewBuffer(req))
	if err != nil {
		log.Println("error in POST method")
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var segment_actual models.Segment
	err = json.NewDecoder(resp.Body).Decode(&segment_actual)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, segment_finish, segment_actual, "Not equal. First test in createSegmentHandler")
}

func TestDeleteSegmentHandler(t *testing.T) {
	testTable := []struct {
		s        SegmentName
		err      error
		response string
	}{
		{s: SegmentName{Name: "AVITO_TEST"}, err: nil, response: "Success"},
		// {s: models.Segment{Id: 2}, err: nil, response: "Success"},
		// {s: models.Segment{Id: 3, Name: "OVITO_TTEST"}, err: nil, response: "Success"},
		{s: SegmentName{}, err: errors.New("Not found"), response: "Error! Not found!"},
	}
	for i, v := range testTable {
		segmentMockDB.EXPECT().Delete(&v.s.Name).Return(v.err)
		req, _ := json.Marshal(v.s.Name)
		resp, err := http.Post("http://localhost:8080/segment/delete", "application/json", bytes.NewBuffer(req))
		if err != nil {
			log.Println("error in POST method. Iteration = ", i)
			log.Fatal(err)
		}
		defer resp.Body.Close()

		var resp_actual string
		err = json.NewDecoder(resp.Body).Decode(&resp_actual)
		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, v.response, resp_actual, fmt.Sprintf("Not equal. Erorr in %d test deleteSegmentHandler", i))

	}
}
