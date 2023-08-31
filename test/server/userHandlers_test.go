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
	"github.com/stretchr/testify/require"
)


func TestFindUserByIdHandler(t *testing.T) {
	excepted := &models.User{
		Id:        1,
		LastName:  "Jones",
		FirstName: "Jon",
		Active:    true,
	}

	userMockDB.EXPECT().FindById(1).Return(excepted, nil)

	m := &Message{1}
	req, _ := json.Marshal(m)
	resp, err := http.Post("http://localhost:8080/user/findByID", "application/json", bytes.NewBuffer(req))
	if err != nil {
		log.Println("tut error ")
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var actual models.User
	err = json.NewDecoder(resp.Body).Decode(&actual)
	if err != nil {
		log.Fatal(err)
	}

	require.Equal(t, *excepted, actual, "Not equal. First test in FindUserById")

	userMockDB.EXPECT().FindById(0).Return(nil, errors.New("uuuupss"))
	m = &Message{0}
	var actual2 string
	req, _ = json.Marshal(m)
	resp, err = http.Post("http://localhost:8080/user/findByID", "application/json", bytes.NewBuffer(req))
	if err != nil {
		log.Println("2 tuta error ")
		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(&actual2)
	if err != nil {
		log.Fatal(err)
	}

	require.Equal(t, "Not found", actual2, "Not Equal. First test in FindUserById")
}

func TestFindAll(t *testing.T) {
	testTable := [][]models.User{
		{
			{
				Id:        1,
				LastName:  "Jones",
				FirstName: "Jon",
				Active:    true,
			},
			{
				Id:        2,
				LastName:  "Morales",
				FirstName: "Miles",
				Active:    true,
			},
			{
				Id:        1,
				LastName:  "Parker",
				FirstName: "Piter",
				Active:    true,
			},
		},
		{},
	}
	for i, v := range testTable {
		userMockDB.EXPECT().FindAll().Return(v, nil)
		resp, err := http.Post("http://localhost:8080/user/findAll", "application/json", nil)
		if err != nil {
			log.Println("FindAll erorr")
			log.Fatal(err)
		}
		defer resp.Body.Close()

		var actual []models.User
		err = json.NewDecoder(resp.Body).Decode(&actual)
		if err != nil {
			log.Fatal(err)
		}

		require.Equal(t, v, actual, fmt.Sprintf("Not equal in %d test in FindAllUsers", i))

	}
}
